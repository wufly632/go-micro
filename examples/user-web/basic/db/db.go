package db

import (
	"database/sql"
	"sync"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/wufly632/go-micro/examples/user-web/basic/config"
)

var (
	inited  bool
	m       sync.RWMutex
	err     error
	mysqlDB *sql.DB
)

func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Info("配置已加载")
		return
	}
	// 判断是否使用mysql
	if config.GetMysqlConfig().GetEnabled() {
		initMysql()
	}
	inited = true
}

func GetDB() *sql.DB {
	return mysqlDB
}
