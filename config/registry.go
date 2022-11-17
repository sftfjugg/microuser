package config

import (
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/nacos/v2"
)

func NewRegistry() registry.Registry {
	r := nacos.NewRegistry(func(options *registry.Options) {
		// nacos注册中心地址
		options.Addrs = []string{"127.0.0.1:8848"}
	})
	return r
}
