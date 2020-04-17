package basic

import (
	"github.com/wufly632/go-micro/examples/user-service/basic/config"
	"github.com/wufly632/go-micro/examples/user-service/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
