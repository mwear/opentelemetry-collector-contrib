module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logzioexporter

go 1.19

require (
	github.com/hashicorp/go-hclog v1.5.0
	github.com/jaegertracing/jaeger v1.41.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.80.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger v0.80.0
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/collector/component v0.89.0
	go.opentelemetry.io/collector/config/configcompression v0.80.1-0.20230629144634-c3f70bd1f8ea
	go.opentelemetry.io/collector/config/confighttp v0.80.1-0.20230629144634-c3f70bd1f8ea
	go.opentelemetry.io/collector/config/configopaque v0.80.1-0.20230629144634-c3f70bd1f8ea
	go.opentelemetry.io/collector/confmap v0.89.0
	go.opentelemetry.io/collector/consumer v0.89.0
	go.opentelemetry.io/collector/exporter v0.89.0
	go.opentelemetry.io/collector/pdata v1.0.0-rcv0018
	go.opentelemetry.io/collector/semconv v0.80.1-0.20230629144634-c3f70bd1f8ea
	go.uber.org/zap v1.26.0
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/apache/thrift v0.18.1 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.16.6 // indirect
	github.com/knadh/koanf v1.5.0 // indirect
	github.com/knadh/koanf/v2 v2.0.1 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20220423185008-bf980b35cac4 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/cors v1.9.0 // indirect
	github.com/uber/jaeger-client-go v2.30.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector v0.89.0 // indirect
	go.opentelemetry.io/collector/config/configauth v0.80.1-0.20230629144634-c3f70bd1f8ea // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.89.0 // indirect
	go.opentelemetry.io/collector/config/configtls v0.80.1-0.20230629144634-c3f70bd1f8ea // indirect
	go.opentelemetry.io/collector/config/internal v0.80.1-0.20230629144634-c3f70bd1f8ea // indirect
	go.opentelemetry.io/collector/extension v0.89.0 // indirect
	go.opentelemetry.io/collector/extension/auth v0.80.1-0.20230629144634-c3f70bd1f8ea // indirect
	go.opentelemetry.io/collector/featuregate v1.0.0-rcv0018 // indirect
	go.opentelemetry.io/collector/receiver v0.89.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.42.0 // indirect
	go.opentelemetry.io/otel v1.20.0 // indirect
	go.opentelemetry.io/otel/metric v1.20.0 // indirect
	go.opentelemetry.io/otel/trace v1.20.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/grpc v1.59.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger => ../../pkg/translator/jaeger

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest
