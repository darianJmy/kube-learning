package config

type Config struct {
	Mysql   *MysqlConfig
	Http    *HttpConfig
}

func NewComponentConfig() *Config {
	options := &Config{
		Mysql: NewMysqlOptions(),
		Http: NewHttpOptions(),

	}
	return options
}