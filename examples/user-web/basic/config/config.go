package config

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/file"
	log "github.com/micro/go-micro/v2/logger"
)

var (
	m                       sync.RWMutex
	inited                  bool                         // 判断是否已初始化过
	sp                      = string(filepath.Separator) //  => '/'
	profiles                defaultProfiles
	etcdConfig              defaultEtcdConfig
	mysqlConfig             defaultMysqlConfig
	err                     error
	defaultRootPath         = "app"
	defaultConfigFilePrefix = "application-"
)

// Init ...
func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Info("[Init] 配置已经初始化过")
		return
	}

	// 加载yml文件
	// Abs返回一个绝对路径
	// filepath.Join(".", sp)  => ./
	// filepath.Dir返回路径
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join(".", sp)))

	// 配置文件路径
	pt := filepath.Join(appPath, "conf")
	os.Chdir(pt)

	// 找到application.yaml文件
	if err = config.Load(file.NewSource(file.WithPath(pt + sp + "application.yml"))); err != nil {
		panic(err)
	}
	// 找到需要引入的新配置文件
	if err = config.Get(defaultRootPath, "profiles").Scan(&profiles); err != nil {
		panic(err)
	}

	log.Infof("[Init] 加载配置文件：path: %s, %+v\n", pt+sp+"application.yml", profiles)

	// 开始导入新文件
	if len(profiles.GetInclude()) > 0 {
		include := strings.Split(profiles.GetInclude(), ",")

		sources := make([]source.Source, len(include))
		for i := 0; i < len(include); i++ {
			filePath := pt + string(filepath.Separator) + defaultConfigFilePrefix + strings.TrimSpace(include[i]) + ".yml"

			log.Infof("[Init] 加载配置文件：path: %s\n", filePath)

			sources[i] = file.NewSource(file.WithPath(filePath))
		}

		// 加载include的文件
		if err = config.Load(sources...); err != nil {
			panic(err)
		}
	}

	// 赋值
	config.Get(defaultRootPath, "etcd").Scan(&etcdConfig)
	config.Get(defaultRootPath, "mysql").Scan(&mysqlConfig)

	// 标记已经初始化
	inited = true

}

func GetMysqlConfig() MysqlConfig {
	return mysqlConfig
}

func GetEtcdConfig() EtcdConfig {
	return etcdConfig
}
