module github.com/douyin-shop/douyin-shop/app/product

go 1.22.9

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/brianvoe/gofakeit/v6 v6.28.0
	github.com/bxcodec/faker/v3 v3.8.1
	github.com/cloudwego/fastpb v0.0.5
	github.com/cloudwego/kitex v0.12.1
	github.com/douyin-shop/douyin-shop/common v0.0.0-20250122041239-0985442f9cfc
	github.com/joho/godotenv v1.5.1
	github.com/kitex-contrib/config-nacos/v2 v2.0.0-20241203120505-63d6739ba3a2
	github.com/kitex-contrib/obs-opentelemetry v0.2.9
	github.com/kitex-contrib/obs-opentelemetry/logging/logrus v0.0.0-20241120035129-55da83caab1b
	github.com/kr/pretty v0.3.1
	github.com/nacos-group/nacos-sdk-go/v2 v2.2.7
	github.com/redis/go-redis/v9 v9.7.0
	go.uber.org/zap v1.27.0
	google.golang.org/protobuf v1.36.3
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	gopkg.in/validator.v2 v2.0.1
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
)

require (
	github.com/alibabacloud-go/debug v1.0.0 // indirect
	github.com/alibabacloud-go/tea v1.1.17 // indirect
	github.com/alibabacloud-go/tea-utils v1.4.4 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1800 // indirect
	github.com/aliyun/alibabacloud-dkms-gcs-go-sdk v0.2.2 // indirect
	github.com/aliyun/alibabacloud-dkms-transfer-go-sdk v0.1.7 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/bytedance/gopkg v0.1.1 // indirect
	github.com/bytedance/sonic v1.12.5 // indirect
	github.com/bytedance/sonic/loader v0.2.0 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/configmanager v0.2.2 // indirect
	github.com/cloudwego/dynamicgo v0.4.7-0.20241220085612-55704ea4ca8f // indirect
	github.com/cloudwego/frugal v0.2.3 // indirect
	github.com/cloudwego/gopkg v0.1.3 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/cloudwego/localsession v0.1.1 // indirect
	github.com/cloudwego/netpoll v0.6.5 // indirect
	github.com/cloudwego/runtimex v0.1.0 // indirect
	github.com/cloudwego/thriftgo v0.3.18 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/go-errors/errors v1.0.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/pprof v0.0.0-20240727154555-813a5fbdbec8 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.25.1 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/jhump/protoreflect v1.8.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kitex-contrib/config-nacos v0.4.1 // indirect
	github.com/kitex-contrib/registry-nacos v0.1.2 // indirect
	github.com/klauspost/cpuid/v2 v2.2.6 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/gls v0.0.0-20220109145502-612d0167dce5 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nacos-group/nacos-sdk-go v1.1.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.18.0 // indirect
	github.com/prometheus/client_model v0.6.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	github.com/sirupsen/logrus v1.9.2 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	github.com/tidwall/gjson v1.17.3 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.45.0 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.20.0 // indirect
	go.opentelemetry.io/contrib/propagators/ot v1.25.0 // indirect
	go.opentelemetry.io/otel v1.34.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.42.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.42.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.34.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.34.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.20.0 // indirect
	go.opentelemetry.io/otel/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.31.0 // indirect
	go.opentelemetry.io/otel/trace v1.34.0 // indirect
	go.opentelemetry.io/proto/otlp v1.5.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/arch v0.6.0 // indirect
	golang.org/x/crypto v0.32.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20240227224415-6ceb2ff114de // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/grpc v1.69.4 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
