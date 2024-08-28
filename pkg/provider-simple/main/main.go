// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"github.com/kubegems/opentofu/pkg/grpcwrap"
	"github.com/kubegems/opentofu/pkg/plugin"
	simple "github.com/kubegems/opentofu/pkg/provider-simple"
	"github.com/kubegems/opentofu/pkg/tfplugin5"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(simple.Provider())
		},
	})
}
