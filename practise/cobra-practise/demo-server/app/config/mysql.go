package config

import (
	"github.com/spf13/pflag"
)

type MysqlConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"name"`
}

func NewMysqlOptions() *MysqlConfig {
	options := &MysqlConfig{}
	return options
}

func (m *MysqlConfig) AddFlags(fs *pflag.FlagSet) {
	if m == nil {
		return
	}
	fs.StringVar(&m.User, "user", m.User, "The media type to use to store user in mysql. ")
	fs.StringVar(&m.Password, "password", m.Password, "The media type to use to password in mysql. ")
	fs.StringVar(&m.Host, "host", m.Host, "The media type to use to store host in mysql. ")
	fs.StringVar(&m.Port, "port", m.Port, "The media type to use to store port in mysql. ")
	fs.StringVar(&m.DBName, "dbname", m.DBName, "The media type to use to store dbname in mysql. ")
}