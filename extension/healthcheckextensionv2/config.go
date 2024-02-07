// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package healthcheckextensionv2 // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextensionv2"

import (
	"errors"
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextensionv2/internal/common"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextensionv2/internal/grpc"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextensionv2/internal/http"
)

var (
	errMissingProtocol      = errors.New("healthcheck extension: must be configured for HTTP or gRPC")
	errGRPCEndpointRequired = errors.New("healthcheck extension: grpc endpoint required")
	errHTTPEndpointRequired = errors.New("healthcheck extension: http endpoint required")
)

// Config has the configuration for the extension enabling the health check
// extension, used to report the health status of the service.
type Config struct {
	ComponentHealthSettings *common.ComponentHealthSettings `mapstructure:"component_health"`
	RecoveryDuration        time.Duration                   `mapstructure:"recovery_duration"`
	GRPCSettings            *grpc.Settings                  `mapstructure:"grpc"`
	HTTPSettings            *http.Settings                  `mapstructure:"http"`
}

func (c *Config) Validate() error {
	if c.GRPCSettings == nil && c.HTTPSettings == nil {
		return errMissingProtocol
	}

	if c.GRPCSettings != nil && c.GRPCSettings.NetAddr.Endpoint == "" {
		return errGRPCEndpointRequired
	}

	if c.HTTPSettings != nil && c.HTTPSettings.Endpoint == "" {
		return errHTTPEndpointRequired
	}

	return nil
}
