module github.com/vincent20101/go-example

go 1.16

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/da440dil/go-workgroup v0.1.0
	github.com/davecgh/go-spew v1.1.1
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-redis/redis/v8 v8.11.3
	github.com/golang/protobuf v1.5.2
	github.com/google/gopacket v1.1.19
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.5.0
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/mitchellh/mapstructure v1.4.3
	github.com/namsral/flag v1.7.4-pre
	github.com/neilotoole/errgroup v0.1.6
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.19.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/vmihailenco/msgpack/v4 v4.2.1
	github.com/vmihailenco/msgpack/v5 v5.3.5
	go.uber.org/goleak v1.1.12 // indirect
	go.uber.org/zap v1.21.0
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/text v0.7.0
	google.golang.org/genproto v0.0.0-20220107163113-42d7afdf6368 // indirect
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	k8s.io/apimachinery v0.24.3
	k8s.io/client-go v0.24.3
	sigs.k8s.io/yaml v1.3.0 // indirect
)

//replace (
//	go-example/grpc_token_auth_test/proto => E:\project\go\go_example\grpc\grpc_token_auth_test/proto
//)
