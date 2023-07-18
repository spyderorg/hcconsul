package telemetry

import (
	"bytes"
	"context"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	gometrics "github.com/armon/go-metrics"
	"github.com/hashicorp/go-hclog"
	"go.opentelemetry.io/otel/attribute"
	otelmetric "go.opentelemetry.io/otel/metric"
	otelsdk "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"

	"github.com/hashicorp/consul/agent/hcp/client"
)

// DefaultExportInterval is a default time interval between export of aggregated metrics.
const DefaultExportInterval = 10 * time.Second

type ConfigProvider interface {
	GetLabels() map[string]string
	GetFilters() *regexp.Regexp
}

// OTELSinkOpts is used to provide configuration when initializing an OTELSink using NewOTELSink.
type OTELSinkOpts struct {
	Reader         otelsdk.Reader
	Ctx            context.Context
	ConfigProvider ConfigProvider
}

// OTELSink captures and aggregates telemetry data as per the OpenTelemetry (OTEL) specification.
// Metric data is exported in OpenTelemetry Protocol (OTLP) wire format.
// This should be used as a Go Metrics backend, as it implements the MetricsSink interface.
type OTELSink struct {
	// spaceReplacer cleans the flattened key by removing any spaces.
	spaceReplacer *strings.Replacer
	logger        hclog.Logger
	cfgProvider   ConfigProvider

	// meterProvider is an OTEL MeterProvider, the entrypoint to the OTEL Metrics SDK.
	// It handles reading/export of aggregated metric data.
	// It enables creation and usage of an OTEL Meter.
	meterProvider *otelsdk.MeterProvider

	// meter is an OTEL Meter, which enables the creation of OTEL instruments.
	meter *otelmetric.Meter

	// Instrument stores contain an OTEL Instrument per metric name (<name, instrument>)
	// for each gauge, counter and histogram types.
	// An instrument allows us to record a measurement for a particular metric, and continuously aggregates metrics.
	// We lazy load the creation of these intruments until a metric is seen, and use them repeatedly to record measurements.
	gaugeInstruments     map[string]otelmetric.Float64ObservableGauge
	counterInstruments   map[string]otelmetric.Float64Counter
	histogramInstruments map[string]otelmetric.Float64Histogram

	// gaugeStore is required to hold last-seen values of gauges
	// This is a workaround, as OTEL currently does not have synchronous gauge instruments.
	// It only allows the registration of "callbacks", which obtain values when the callback is called.
	// We must hold gauge values until the callback is called, when the measurement is exported, and can be removed.
	gaugeStore *gaugeStore

	mutex sync.Mutex
}

// NewOTELReader returns a configured OTEL PeriodicReader to export metrics every X seconds.
// It configures the reader with a custom OTELExporter with a MetricsClient to transform and export
// metrics in OTLP format to an external url.
func NewOTELReader(client client.MetricsClient, endpointProvider EndpointProvider, exportInterval time.Duration) otelsdk.Reader {
	exporter := NewOTELExporter(client, endpointProvider)
	return otelsdk.NewPeriodicReader(exporter, otelsdk.WithInterval(exportInterval))
}

// NewOTELSink returns a sink which fits the Go Metrics MetricsSink interface.
// It sets up a MeterProvider and Meter, key pieces of the OTEL Metrics SDK which
// enable us to create OTEL Instruments to record measurements.
func NewOTELSink(opts *OTELSinkOpts) (*OTELSink, error) {
	if opts.Reader == nil {
		return nil, fmt.Errorf("ferror: provide valid reader")
	}

	if opts.Ctx == nil {
		return nil, fmt.Errorf("ferror: provide valid context")
	}

	if opts.ConfigProvider == nil {
		return nil, fmt.Errorf("ferror: provide valid config provider")
	}

	logger := hclog.FromContext(opts.Ctx).Named("otel_sink")

	// Setup OTEL Metrics SDK to aggregate, convert and export metrics periodically.
	res := resource.NewSchemaless()
	meterProvider := otelsdk.NewMeterProvider(otelsdk.WithResource(res), otelsdk.WithReader(opts.Reader))
	meter := meterProvider.Meter("github.com/hashicorp/consul/agent/hcp/telemetry")

	return &OTELSink{
		cfgProvider:          opts.ConfigProvider,
		spaceReplacer:        strings.NewReplacer(" ", "_"),
		logger:               logger,
		meterProvider:        meterProvider,
		meter:                &meter,
		gaugeStore:           NewGaugeStore(),
		gaugeInstruments:     make(map[string]otelmetric.Float64ObservableGauge, 0),
		counterInstruments:   make(map[string]otelmetric.Float64Counter, 0),
		histogramInstruments: make(map[string]otelmetric.Float64Histogram, 0),
	}, nil
}

// SetGauge emits a Consul gauge metric.
func (o *OTELSink) SetGauge(key []string, val float32) {
	o.SetGaugeWithLabels(key, val, nil)
}

// AddSample emits a Consul histogram metric.
func (o *OTELSink) AddSample(key []string, val float32) {
	o.AddSampleWithLabels(key, val, nil)
}

// IncrCounter emits a Consul counter metric.
func (o *OTELSink) IncrCounter(key []string, val float32) {
	o.IncrCounterWithLabels(key, val, nil)
}

// AddSampleWithLabels emits a Consul gauge metric that gets
// registed by an OpenTelemetry Histogram instrument.
func (o *OTELSink) SetGaugeWithLabels(key []string, val float32, labels []gometrics.Label) {
	k := o.flattenKey(key)

	if !o.cfgProvider.GetFilters().MatchString(k) {
		return
	}

	// Set value in global Gauge store.
	o.gaugeStore.Set(k, float64(val), o.labelsToAttributes(labels))

	o.mutex.Lock()
	defer o.mutex.Unlock()

	// If instrument does not exist, create it and register callback to emit last value in global Gauge store.
	if _, ok := o.gaugeInstruments[k]; !ok {
		// The registration of a callback only needs to happen once, when the instrument is created.
		// The callback will be triggered every export cycle for that metric.
		// It must be explicitly de-registered to be removed (which we do not do), to ensure new gauge values are exported every cycle.
		inst, err := (*o.meter).Float64ObservableGauge(k, otelmetric.WithFloat64Callback(o.gaugeStore.gaugeCallback(k)))
		if err != nil {
			o.logger.Error("Failed to create gauge instrument", "error", err)
			return
		}
		o.gaugeInstruments[k] = inst
	}
}

// AddSampleWithLabels emits a Consul sample metric that gets registed by an OpenTelemetry Histogram instrument.
func (o *OTELSink) AddSampleWithLabels(key []string, val float32, labels []gometrics.Label) {
	k := o.flattenKey(key)

	if !o.cfgProvider.GetFilters().MatchString(k) {
		return
	}

	o.mutex.Lock()
	defer o.mutex.Unlock()

	inst, ok := o.histogramInstruments[k]
	if !ok {
		histogram, err := (*o.meter).Float64Histogram(k)
		if err != nil {
			o.logger.Error("Failed create histogram instrument", "error", err)
			return
		}
		inst = histogram
		o.histogramInstruments[k] = inst
	}

	attrs := o.labelsToAttributes(labels)
	inst.Record(context.TODO(), float64(val), otelmetric.WithAttributes(attrs...))
}

// IncrCounterWithLabels emits a Consul counter metric that gets registed by an OpenTelemetry Histogram instrument.
func (o *OTELSink) IncrCounterWithLabels(key []string, val float32, labels []gometrics.Label) {
	k := o.flattenKey(key)

	if !o.cfgProvider.GetFilters().MatchString(k) {
		return
	}

	o.mutex.Lock()
	defer o.mutex.Unlock()

	inst, ok := o.counterInstruments[k]
	if !ok {
		counter, err := (*o.meter).Float64Counter(k)
		if err != nil {
			o.logger.Error("Failed to create counter instrument:", "error", err)
			return
		}

		inst = counter
		o.counterInstruments[k] = inst
	}

	attrs := o.labelsToAttributes(labels)
	inst.Add(context.TODO(), float64(val), otelmetric.WithAttributes(attrs...))
}

// EmitKey unsupported.
func (o *OTELSink) EmitKey(key []string, val float32) {}

// flattenKey key along with its labels.
func (o *OTELSink) flattenKey(parts []string) string {
	buf := &bytes.Buffer{}
	joined := strings.Join(parts, ".")

	o.spaceReplacer.WriteString(buf, joined)

	return buf.String()
}

// labelsToAttributes converts go metrics and provider labels into OTEL format []attributes.KeyValue
func (o *OTELSink) labelsToAttributes(goMetricsLabels []gometrics.Label) []attribute.KeyValue {
	providerLabels := o.cfgProvider.GetLabels()

	length := len(goMetricsLabels) + len(providerLabels)
	if length == 0 {
		return []attribute.KeyValue{}
	}

	attrs := make([]attribute.KeyValue, 0, length)
	// Convert provider labels to OTEL attributes.
	for _, label := range goMetricsLabels {
		attrs = append(attrs, attribute.KeyValue{
			Key:   attribute.Key(label.Name),
			Value: attribute.StringValue(label.Value),
		})
	}

	// Convert provider labels to OTEL attributes.
	for k, v := range providerLabels {
		attrs = append(attrs, attribute.KeyValue{
			Key:   attribute.Key(k),
			Value: attribute.StringValue(v),
		})
	}

	return attrs
}