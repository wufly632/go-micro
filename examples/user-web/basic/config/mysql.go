package config

import "time"

type defaultMysqlConfig struct {
	URL               string        `json:"url"`
	Enable            bool          `json:"enabled"`
	MaxIdleConnection int           `json:"maxIdleConnection"`
	MaxOpenConnection int           `json:"maxOpenConnection"`
	ConnMaxLifetime   time.Duration `json:"connMaxLifetime"`
}

type MysqlConfig interface {
	GetURL() string
	GetEnabled() bool
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
	GetConnMaxLifetime() time.Duration
}

func (d defaultMysqlConfig) GetURL() string {
	return d.URL
}

func (d defaultMysqlConfig) GetEnabled() bool {
	return d.Enable
}

func (d defaultMysqlConfig) GetMaxIdleConnection() int {
	return d.MaxIdleConnection
}

func (d defaultMysqlConfig) GetMaxOpenConnection() int {
	return d.MaxOpenConnection
}

func (d defaultMysqlConfig) GetConnMaxLifetime() time.Duration {
	return d.ConnMaxLifetime
}
