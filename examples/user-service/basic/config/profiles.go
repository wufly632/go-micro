package config

// defaultProfiles 属性配置文件
type defaultProfiles struct {
	Include string `json:"include"`
}

func (p defaultProfiles) GetInclude() string {
	return p.Include
}
