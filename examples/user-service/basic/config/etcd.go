package config

type defaultEtcdConfig struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

type EtcdConfig interface {
	GetEnabled() bool
	GetPort() int
	GetHost() string
}

// GetPort Etcd 端口
func (c defaultEtcdConfig) GetPort() int {
	return c.Port
}

// GetEnabled Etcd 激活
func (c defaultEtcdConfig) GetEnabled() bool {
	return c.Enabled
}

// GetHost Etcd 主机地址
func (c defaultEtcdConfig) GetHost() string {
	return c.Host
}
