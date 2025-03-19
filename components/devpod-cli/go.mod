module github.com/khulnasoft/devpod/devpod-cli

go 1.22.0

toolchain go1.23.3

require (
	github.com/bufbuild/connect-go v1.10.0
	github.com/creack/pty v1.1.17
	github.com/khulnasoft/devpod/common-go v0.0.0-00010101000000-000000000000
	github.com/khulnasoft/devpod/components/public-api/go v0.0.0-20230220133850-852f5cd5b180
	github.com/khulnasoft/devpod/devpod-protocol v0.0.0-00010101000000-000000000000
	github.com/khulnasoft/devpod/ide-metrics-api v0.0.0-00010101000000-000000000000
	github.com/khulnasoft/devpod/supervisor/api v0.0.0-00010101000000-000000000000
	github.com/go-errors/errors v1.4.2
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.6.0
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510
	github.com/google/tcpproxy v0.0.0-20180808230851-dfa16c61dad2
	github.com/gorilla/handlers v1.5.1
	github.com/manifoldco/promptui v0.9.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/prometheus/procfs v0.10.1
	github.com/sirupsen/logrus v1.9.3
	github.com/sourcegraph/jsonrpc2 v0.0.0-20200429184054-15c2290dcb37
	github.com/spf13/cobra v1.7.0
	github.com/stretchr/testify v1.8.4
	golang.org/x/sync v0.5.0
	golang.org/x/term v0.18.0
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2
	google.golang.org/grpc v1.58.3
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/khulnasoft/devpod/components/scrubber v0.0.0-00010101000000-000000000000 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/mattn/go-colorable v0.0.9 // indirect
	github.com/mattn/go-isatty v0.0.3 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.29.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.16.0 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	google.golang.org/genproto v0.0.0-20230803162519-f966b187b2e5 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230726155614-23370e0ffb3e // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/segmentio/backo-go v0.0.0-20200129164019-23eae7c10bd3 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/x-cray/logrus-prefixed-formatter v0.5.2
	github.com/xtgo/uuid v0.0.0-20140804021211-a0b114877d4c // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/segmentio/analytics-go.v3 v3.1.0 // indirect
)

replace github.com/khulnasoft/devpod/common-go => ../common-go // leeway

replace github.com/khulnasoft/devpod/components/public-api/go => ../public-api/go // leeway

replace github.com/khulnasoft/devpod/components/scrubber => ../scrubber // leeway

replace github.com/khulnasoft/devpod/devpod-protocol => ../devpod-protocol/go // leeway

replace github.com/khulnasoft/devpod/ide-metrics-api => ../ide-metrics-api/go // leeway

replace github.com/khulnasoft/devpod/supervisor/api => ../supervisor-api/go // leeway

replace github.com/google/addlicense => ../../dev/addlicense // leeway

replace k8s.io/api => k8s.io/api v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/apimachinery => k8s.io/apimachinery v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/apiserver => k8s.io/apiserver v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/client-go => k8s.io/client-go v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/code-generator => k8s.io/code-generator v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/component-base => k8s.io/component-base v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/cri-api => k8s.io/cri-api v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/kubelet => k8s.io/kubelet v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/metrics => k8s.io/metrics v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/component-helpers => k8s.io/component-helpers v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/controller-manager => k8s.io/controller-manager v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/kubectl => k8s.io/kubectl v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/mount-utils => k8s.io/mount-utils v0.30.9 // leeway indirect from components/common-go:lib

replace k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.30.9 // leeway indirect from components/common-go:lib
