// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testhelpers

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextensionv2/internal/status"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"gopkg.in/yaml.v3"
)

type PipelineMetadata struct {
	PipelineID  component.ID
	ReceiverID  *component.InstanceID
	ProcessorID *component.InstanceID
	ExporterID  *component.InstanceID
}

func (p *PipelineMetadata) InstanceIDs() []*component.InstanceID {
	return []*component.InstanceID{p.ReceiverID, p.ProcessorID, p.ExporterID}
}

func NewPipelineMetadata(typeVal component.Type) *PipelineMetadata {
	pipelineID := component.NewID(typeVal)
	return &PipelineMetadata{
		PipelineID: pipelineID,
		ReceiverID: &component.InstanceID{
			ID:   component.NewIDWithName(typeVal, "in"),
			Kind: component.KindReceiver,
			PipelineIDs: map[component.ID]struct{}{
				pipelineID: {},
			},
		},
		ProcessorID: &component.InstanceID{
			ID:   component.NewID("batch"),
			Kind: component.KindProcessor,
			PipelineIDs: map[component.ID]struct{}{
				pipelineID: {},
			},
		},
		ExporterID: &component.InstanceID{
			ID:   component.NewIDWithName(typeVal, "out"),
			Kind: component.KindExporter,
			PipelineIDs: map[component.ID]struct{}{
				pipelineID: {},
			},
		},
	}
}

func NewPipelines(typeVals ...component.Type) map[string]*PipelineMetadata {
	result := make(map[string]*PipelineMetadata, len(typeVals))
	for _, val := range typeVals {
		result[string(val)] = NewPipelineMetadata(val)
	}
	return result
}

// SeedAggregator records a status event for each instanceID
func SeedAggregator(
	agg *status.Aggregator,
	instanceIDs []*component.InstanceID,
	statuses ...component.Status,
) {
	for _, st := range statuses {
		for _, id := range instanceIDs {
			agg.RecordStatus(id, component.NewStatusEvent(st))
		}
	}
}

func NewConfmapFromFile(t *testing.T, fileName string) (*confmap.Conf, error) {
	content, err := os.ReadFile(filepath.Clean(fileName))
	if err != nil {
		return nil, err
	}

	var data map[string]any
	if err = yaml.Unmarshal(content, &data); err != nil {
		return nil, err
	}

	return confmap.NewFromStringMap(data), nil
}
