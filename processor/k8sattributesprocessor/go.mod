module github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor

go 1.26.0

require (
	github.com/google/go-cmp v0.7.0
	github.com/google/uuid v1.6.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig v0.80.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8stest v0.80.0
	github.com/stretchr/testify v1.12.1
	go.opencensus.io v0.24.0
	go.opentelemetry.io/collector v0.80.1-0.20230629144634-c3f70bd1f8ea
	go.opentelemetry.io/collector/component v1.69.0
	go.opentelemetry.io/collector/component/componenttest v0.163.0
	go.opentelemetry.io/collector/confmap v1.69.0
	go.opentelemetry.io/collector/consumer v1.69.0
	go.opentelemetry.io/collector/consumer/consumertest v0.163.0
	go.opentelemetry.io/collector/pdata v1.69.0
	go.opentelemetry.io/collector/processor v0.80.1-0.20230629144634-c3f70bd1f8ea
	go.opentelemetry.io/collector/receiver/otlpreceiver v0.163.0
	go.opentelemetry.io/collector/receiver/receivertest v0.163.0
	go.opentelemetry.io/collector/semconv v0.80.1-0.20230629144634-c3f70bd1f8ea
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.28.0
	k8s.io/api v0.27.3
	k8s.io/apimachinery v0.27.3
	k8s.io/client-go v0.27.3
)

require (
	github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
	github.com/Microsoft/go-winio v0.6.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/distribution v2.8.2+incompatible // indirect
	github.com/docker/docker v24.0.2+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/emicklei/go-restful/v3 v3.9.0 // indirect
	github.com/evanphx/json-patch v5.6.0+incompatible // indirect
	github.com/felixge/httpsnoop v1.1.0 // indirect
	github.com/foxboron/go-tpm-keyfiles v0.0.0-20251226215517-609e4778396f // indirect
	github.com/go-logr/logr v1.4.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.1 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-viper/mapstructure/v2 v2.5.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v1.0.0 // indirect
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/google/go-tpm v0.9.8 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/hashicorp/go-version v1.9.0 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.20.3 // indirect
	github.com/knadh/koanf v1.5.0 // indirect
	github.com/knadh/koanf/v2 v2.3.8 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.3-0.20250322232337-35a7c28c31ee // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/openshift/api v3.9.0+incompatible // indirect
	github.com/openshift/client-go v0.0.0-20210521082421-73d9475a9142 // indirect
	github.com/pierrec/lz4/v4 v4.1.33 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/collector/component/componentstatus v0.163.0 // indirect
	go.opentelemetry.io/collector/config/configauth v1.69.0 // indirect
	go.opentelemetry.io/collector/config/configcompression v1.69.0 // indirect
	go.opentelemetry.io/collector/config/configgrpc v1.69.0 // indirect
	go.opentelemetry.io/collector/config/confighttp v0.163.0 // indirect
	go.opentelemetry.io/collector/config/configmiddleware v1.69.0 // indirect
	go.opentelemetry.io/collector/config/confignet v1.69.0 // indirect
	go.opentelemetry.io/collector/config/configopaque v1.69.0 // indirect
	go.opentelemetry.io/collector/config/configoptional v1.69.0 // indirect
	go.opentelemetry.io/collector/config/configtls v1.69.0 // indirect
	go.opentelemetry.io/collector/consumer/consumererror v0.163.0 // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.163.0 // indirect
	go.opentelemetry.io/collector/extension/extensionauth v1.69.0 // indirect
	go.opentelemetry.io/collector/extension/extensionmiddleware v0.163.0 // indirect
	go.opentelemetry.io/collector/featuregate v1.69.0 // indirect
	go.opentelemetry.io/collector/internal/componentalias v0.163.0 // indirect
	go.opentelemetry.io/collector/internal/statusutil v0.163.0 // indirect
	go.opentelemetry.io/collector/internal/telemetry v0.163.0 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.163.0 // indirect
	go.opentelemetry.io/collector/pipeline v1.69.0 // indirect
	go.opentelemetry.io/collector/pipeline/xpipeline v0.163.0 // indirect
	go.opentelemetry.io/collector/receiver v1.69.0 // indirect
	go.opentelemetry.io/collector/receiver/receiverhelper v0.163.0 // indirect
	go.opentelemetry.io/collector/receiver/xreceiver v0.163.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.70.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.70.0 // indirect
	go.opentelemetry.io/otel v1.47.0 // indirect
	go.opentelemetry.io/otel/log v1.47.0 // indirect
	go.opentelemetry.io/otel/metric v1.47.0 // indirect
	go.opentelemetry.io/otel/sdk v1.47.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.47.0 // indirect
	go.opentelemetry.io/otel/trace v1.47.0 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/crypto v0.58.0 // indirect
	golang.org/x/net v0.61.0 // indirect
	golang.org/x/oauth2 v0.36.0 // indirect
	golang.org/x/sys v0.49.0 // indirect
	golang.org/x/term v0.47.0 // indirect
	golang.org/x/text v0.43.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260803160001-6ac0973c030d // indirect
	google.golang.org/grpc v1.84.0 // indirect
	google.golang.org/protobuf v1.36.12 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/klog/v2 v2.90.1 // indirect
	k8s.io/kube-openapi v0.0.0-20230501164219-8b0f38b5fd1f // indirect
	k8s.io/utils v0.0.0-20230209194617-a36077c30491 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig => ./../../internal/k8sconfig

// openshift removed all tags from their repo, use the pseudoversion from the release-3.9 branch HEAD
replace github.com/openshift/api v3.9.0+incompatible => github.com/openshift/api v0.0.0-20180801171038-322a19404e37

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8stest => ../../internal/k8stest

// ambiguous import: found package cloud.google.com/go/compute/metadata in multiple modules
replace cloud.google.com/go v0.54.0 => cloud.google.com/go v0.110.2
