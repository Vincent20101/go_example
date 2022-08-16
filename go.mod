module github.com/vincent20101/go-example

go 1.16

require (
	github.com/da440dil/go-workgroup v0.1.0
	github.com/davecgh/go-spew v1.1.1
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/websocket v1.5.0
	github.com/mitchellh/mapstructure v1.4.3
	github.com/neilotoole/errgroup v0.1.6
	github.com/pkg/errors v0.9.1
	go.uber.org/zap v1.21.0
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
	k8s.io/apimachinery v0.24.3
	k8s.io/client-go v0.24.3
)

//replace (
//	go-example/grpc_token_auth_test/proto => E:\project\go\go_example\grpc\grpc_token_auth_test/proto
//)
