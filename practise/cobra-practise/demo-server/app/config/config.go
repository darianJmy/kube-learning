package config

type Config struct {
	Mysql   *MysqlConfig
}

func NewComponentConfig() *Config {
	options := &Config{
		Mysql: NewMysqlOptions(),
	}
	return options
}