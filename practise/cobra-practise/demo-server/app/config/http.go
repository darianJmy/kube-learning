package config

import (
	"github.com/spf13/pflag"
)

type HttpConfig struct {
	HttpPort string `yaml:"httpport"`
	HttpUrl  string `yaml:"httpurl"`
}

func NewHttpOptions() *HttpConfig {
	options := &HttpConfig{}
	return options
}

func (h *HttpConfig) AddFlags(fs *pflag.FlagSet) {
	if h == nil {
		return
	}
	fs.StringVar(&h.HttpPort, "httpport", h.HttpPort, "The media type to use to store httpport in http. ")
	fs.StringVar(&h.HttpUrl, "httpurl", h.HttpUrl, "The media type to use to httpport in mysql. ")
}
