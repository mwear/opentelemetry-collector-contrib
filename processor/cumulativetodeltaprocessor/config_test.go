// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cumulativetodeltaprocessor

import (
	"path/filepath"
	"testing"
	"time"

	"go.opentelemetry.io/collector/confmap/confmaptest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/processor/filterset"
)

func TestLoadConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		id           config.ComponentID
		expected     config.Processor
		errorMessage string
	}{
		{
			id: config.NewComponentIDWithName(typeStr, ""),
			expected: &Config{
				ProcessorSettings: config.NewProcessorSettings(config.NewComponentID(typeStr)),
				Include: MatchMetrics{
					Metrics: []string{
						"metric1",
						"metric2",
					},
					Config: filterset.Config{
						MatchType:    "strict",
						RegexpConfig: nil,
					},
				},
				Exclude: MatchMetrics{
					Metrics: []string{
						"metric3",
						"metric4",
					},
					Config: filterset.Config{
						MatchType:    "strict",
						RegexpConfig: nil,
					},
				},
				MaxStaleness: 10 * time.Second,
			},
		},
		{
			id:       config.NewComponentIDWithName(typeStr, "empty"),
			expected: createDefaultConfig(),
		},
		{
			id: config.NewComponentIDWithName(typeStr, "regexp"),
			expected: &Config{
				ProcessorSettings: config.NewProcessorSettings(config.NewComponentID(typeStr)),
				Include: MatchMetrics{
					Metrics: []string{
						"a*",
					},
					Config: filterset.Config{
						MatchType:    "regexp",
						RegexpConfig: nil,
					},
				},
				Exclude: MatchMetrics{
					Metrics: []string{
						"b*",
					},
					Config: filterset.Config{
						MatchType:    "regexp",
						RegexpConfig: nil,
					},
				},
				MaxStaleness: 10 * time.Second,
			},
		},
		{
			id:           config.NewComponentIDWithName(typeStr, "missing_match_type"),
			errorMessage: "match_type must be set if metrics are supplied",
		},
		{
			id:           config.NewComponentIDWithName(typeStr, "missing_name"),
			errorMessage: "metrics must be supplied if match_type is set",
		},
	}

	for _, tt := range tests {
		t.Run(tt.id.String(), func(t *testing.T) {
			cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
			require.NoError(t, err)

			factory := NewFactory()
			cfg := factory.CreateDefaultConfig()

			sub, err := cm.Sub(tt.id.String())
			require.NoError(t, err)
			require.NoError(t, config.UnmarshalProcessor(sub, cfg))

			if tt.expected == nil {
				assert.EqualError(t, cfg.Validate(), tt.errorMessage)
				return
			}
			assert.NoError(t, cfg.Validate())
			assert.Equal(t, tt.expected, cfg)
		})
	}
}
