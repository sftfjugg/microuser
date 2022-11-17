module github.com/sftfjugg/microuser

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/nacos/v2 v2.0.0-20201025091542-fa097e59f8ac
	github.com/nacos-group/nacos-sdk-go v1.0.9
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.2.4
)
