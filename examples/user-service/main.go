package main

import (
	"fmt"
	"time"

	"github.com/wufly632/go-micro/examples/user-service/basic/config"
	"github.com/wufly632/go-micro/examples/user-service/model"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/wufly632/go-micro/examples/user-service/basic"
	"github.com/wufly632/go-micro/examples/user-service/handler"

	user "github.com/wufly632/go-micro/examples/user-service/proto/user"
)

func main() {
	// 启动配置文件
	basic.Init()

	// 使用etch注册中心
	micReg := etcd.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name("wufly.micro.book.service.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) error {
			// 初始化模型
			model.Init()
			// 初始化handler
			handler.Init()
			return nil
		}),
	)

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.Service))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
