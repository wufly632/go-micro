package main

import (
	"fmt"
	"time"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/cli/v2"
	"github.com/wufly632/go-micro/examples/user-web/basic"
	"github.com/wufly632/go-micro/examples/user-web/basic/config"
	"github.com/wufly632/go-micro/examples/user-web/handler"
)

func main() {
	basic.Init()
	// 使用etcd注册中心
	micReg := etcd.NewRegistry(registryOptions)
	// create new web service
	service := web.NewService(
		web.Name("wufly.micro.book.web.user"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8088"),
	)

	// initialise service
	if err := service.Init(
		web.Action(func(c *cli.Context) {
			// 初始化handler
			handler.Init()
		}),
	); err != nil {
		log.Fatal(err)
	}

	// register call handler
	service.HandleFunc("/user/login", handler.Login)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
