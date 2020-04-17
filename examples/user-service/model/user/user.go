package user

import (
	"fmt"
	"sync"

	proto "github.com/wufly632/go-micro/examples/user-service/proto/user"
)

type service struct {
}

type Service interface {
	QueryUserByName(name string) (ret *proto.User, err error)
}

var (
	s Service
	m sync.RWMutex
)

func GetService() (Service, error) {
	if s == nil {
		return s, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init ...
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}
	s = &service{}
}
