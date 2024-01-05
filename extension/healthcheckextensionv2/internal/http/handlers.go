// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"encoding/json"
	"net/http"

	"go.opentelemetry.io/collector/component"
)

var responseCodes = map[component.Status]int{
	component.StatusNone:             http.StatusServiceUnavailable,
	component.StatusStarting:         http.StatusServiceUnavailable,
	component.StatusOK:               http.StatusOK,
	component.StatusRecoverableError: http.StatusServiceUnavailable,
	component.StatusPermanentError:   http.StatusBadRequest,
	component.StatusFatalError:       http.StatusInternalServerError,
	component.StatusStopping:         http.StatusServiceUnavailable,
	component.StatusStopped:          http.StatusServiceUnavailable,
}

func (s *Server) statusHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var sst *serializableStatus
		pipeline := r.URL.Query().Get("pipeline")

		if pipeline == "" {
			details := s.aggregator.CollectorStatusDetailed()
			sst = toCollectorSerializableStatus(details)
		} else {
			details, err := s.aggregator.PipelineStatusDetailed(pipeline)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			sst = toPipelineSerializableStatus(details)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseCodes[sst.Status()])

		body, _ := json.Marshal(sst)
		_, _ = w.Write(body)
	})
}

func (s *Server) configHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		conf := func() []byte {
			s.mu.RLock()
			defer s.mu.RUnlock()
			return s.colconf
		}()

		if len(conf) == 0 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(conf)
	})
}
