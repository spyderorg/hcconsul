// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types

import (
	"google.golang.org/protobuf/proto"

	pbmesh "github.com/hashicorp/consul/proto-public/pbmesh/v1alpha1"
)

type XRouteData interface {
	proto.Message
	XRouteWithRefs
}

type XRouteWithRefs interface {
	GetParentRefs() []*pbmesh.ParentReference
	GetUnderlyingBackendRefs() []*pbmesh.BackendReference
}