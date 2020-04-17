package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/wufly632/go-micro/examples/user-web/basic/config"
)

func initMysql() {
	// 创建连接
	mysqlDB, err = sql.Open("mysql", config.GetMysqlConfig().GetURL())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	// 最大连接数
	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())
	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())
	mysqlDB.SetConnMaxLifetime(time.Second * time.Duration(config.GetMysqlConfig().GetConnMaxLifetime()))
	// 测试连接
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
}
