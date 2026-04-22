module github.com/khulnasoft/devpod/components/devpod-db/go

go 1.24.0

replace github.com/khulnasoft/devpod/common-go => ../../common-go // blazedock

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.6.0
	github.com/khulnasoft/devpod/common-go v0.0.0-00010101000000-000000000000
	github.com/relvacode/iso8601 v1.1.0
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.11.1
	google.golang.org/grpc v1.79.3
	google.golang.org/protobuf v1.36.10
	gorm.io/datatypes v1.0.7
	gorm.io/driver/mysql v1.4.4
	gorm.io/gorm v1.25.1
	gorm.io/plugin/opentelemetry v0.1.3
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/khulnasoft/devpod/components/scrubber v0.0.0-00010101000000-000000000000 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.16.0 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.10.1 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/otel v1.39.0 // indirect
	go.opentelemetry.io/otel/metric v1.39.0 // indirect
	go.opentelemetry.io/otel/trace v1.39.0 // indirect
	golang.org/x/net v0.48.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/khulnasoft/devpod/components/scrubber => ../../scrubber // blazedock
