package grpc

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/collector/component"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

func (s *Server) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	var err error
	var ev *component.StatusEvent

	if req.Service == "" {
		ev = s.aggregator.CollectorStatus()
	} else {
		ev, err = s.aggregator.PipelineStatus(req.Service)
	}

	if err != nil {
		return nil, status.Error(codes.NotFound, "unknown service")
	}

	return &healthpb.HealthCheckResponse{
		Status: s.toServingStatus(ev),
	}, nil
}

func (s *Server) Watch(req *healthpb.HealthCheckRequest, stream healthpb.Health_WatchServer) error {
	sub, err := s.aggregator.Subscribe(req.Service)
	if err != nil {
		return err
	}
	defer s.aggregator.Unsubscribe(sub)

	var lastServingStatus healthpb.HealthCheckResponse_ServingStatus = -1

	failureTicker := time.NewTicker(s.failureDuration)
	failureTicker.Stop()

	for {
		select {
		case ev, ok := <-sub:
			if !ok {
				return status.Error(codes.Canceled, "Server shutting down.")
			}

			var sst healthpb.HealthCheckResponse_ServingStatus

			switch {
			case ev == nil:
				sst = healthpb.HealthCheckResponse_SERVICE_UNKNOWN
			case ev.Status() == component.StatusRecoverableError:
				fmt.Printf("recoverable error: setting timer: %s\n", ev.Err().Error())
				failureTicker.Reset(s.failureDuration)
				sst = lastServingStatus
				if lastServingStatus == -1 {
					sst = healthpb.HealthCheckResponse_SERVING
				}
			default:
				failureTicker.Stop()
				sst = statusToServingStatusMap[ev.Status()]
				fmt.Printf("setting sst: %s, evs: %s\n", sst, ev.Status().String())
			}

			if lastServingStatus == sst {
				fmt.Printf("skipping status same: %s\n", sst)
				continue
			}

			lastServingStatus = sst

			err := stream.Send(&healthpb.HealthCheckResponse{Status: sst})
			if err != nil {
				return status.Error(codes.Canceled, "Stream has ended.")
			}
		case <-failureTicker.C:
			fmt.Println("failure ticker triggered")
			failureTicker.Stop()
			if lastServingStatus == healthpb.HealthCheckResponse_NOT_SERVING {
				continue
			}
			lastServingStatus = healthpb.HealthCheckResponse_NOT_SERVING
			err := stream.Send(
				&healthpb.HealthCheckResponse{
					Status: healthpb.HealthCheckResponse_NOT_SERVING,
				},
			)
			if err != nil {
				return status.Error(codes.Canceled, "Stream has ended.")
			}
		case <-stream.Context().Done():
			return status.Error(codes.Canceled, "Stream has ended.")
		}
	}
}

var statusToServingStatusMap = map[component.Status]healthpb.HealthCheckResponse_ServingStatus{
	component.StatusNone:             healthpb.HealthCheckResponse_NOT_SERVING,
	component.StatusStarting:         healthpb.HealthCheckResponse_NOT_SERVING,
	component.StatusOK:               healthpb.HealthCheckResponse_SERVING,
	component.StatusRecoverableError: healthpb.HealthCheckResponse_SERVING,
	component.StatusPermanentError:   healthpb.HealthCheckResponse_NOT_SERVING,
	component.StatusFatalError:       healthpb.HealthCheckResponse_NOT_SERVING,
	component.StatusStopping:         healthpb.HealthCheckResponse_NOT_SERVING,
	component.StatusStopped:          healthpb.HealthCheckResponse_NOT_SERVING,
}

func (s *Server) toServingStatus(ev *component.StatusEvent) healthpb.HealthCheckResponse_ServingStatus {
	if ev.Status() == component.StatusRecoverableError &&
		time.Now().Compare(ev.Timestamp().Add(s.failureDuration)) == 1 {
		return healthpb.HealthCheckResponse_NOT_SERVING
	}
	return statusToServingStatusMap[ev.Status()]
}
