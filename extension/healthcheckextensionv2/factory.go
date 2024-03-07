// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate mdatagen metadata.yaml

package healthcheckextensionv2 // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextensionv2"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/extension"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextensionv2/internal/grpc"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextensionv2/internal/http"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextensionv2/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/localhostgate"
)

const (
	defaultGRPCPort = 13132
	defaultHTTPPort = 13133
)

// NewFactory creates a factory for HealthCheck extension.
func NewFactory() extension.Factory {
	return extension.NewFactory(
		metadata.Type,
		createDefaultConfig,
		createExtension,
		metadata.ExtensionStability,
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		LegacyConfig: http.LegacyConfig{
			ServerConfig: confighttp.ServerConfig{
				Endpoint: localhostgate.EndpointForPort(defaultHTTPPort),
			},
			Path: "/",
		},
		HTTPConfig: &http.Config{
			ServerConfig: confighttp.ServerConfig{
				Endpoint: localhostgate.EndpointForPort(defaultHTTPPort),
			},
			Status: http.PathConfig{
				Enabled: true,
				Path:    "/status",
			},
			Config: http.PathConfig{
				Enabled: false,
				Path:    "/config",
			},
		},
		GRPCConfig: &grpc.Config{
			ServerConfig: configgrpc.ServerConfig{
				NetAddr: confignet.AddrConfig{
					Endpoint:  localhostgate.EndpointForPort(defaultGRPCPort),
					Transport: "tcp",
				},
			},
		},
	}
}

func createExtension(ctx context.Context, set extension.CreateSettings, cfg component.Config) (extension.Extension, error) {
	config := cfg.(*Config)
	return newExtension(ctx, *config, set), nil
}
