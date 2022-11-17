package main

import (
	"fmt"
	"github.com/sftfjugg/microuser/config"
	"github.com/sftfjugg/microuser/handler"
	microuser "github.com/sftfjugg/microuser/proto/microuser"
	"github.com/sftfjugg/microuser/subscriber"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	// 使用naocs配置中心
	config.InitSetting()
	// 获取nacos注册中心实例
	r := config.NewRegistry()

	// New Service
	service := micro.NewService(
		micro.Name(config.GlobalConf.App.Name),
		micro.Address(fmt.Sprintf(":%d", config.GlobalConf.App.Port)),
		micro.Version("latest"),
		micro.Registry(r),
	)

	// Initialise service
	service.Init()

	// Register Handler
	microuser.RegisterMicrouserHandler(service.Server(), new(handler.Microuser))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.microuser", service.Server(), new(subscriber.Microuser))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
