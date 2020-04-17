package basic

import (
	"github.com/wufly632/go-micro/examples/user-web/basic/config"
	"github.com/wufly632/go-micro/examples/user-web/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
