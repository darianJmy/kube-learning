package options

import (
	"github.com/spf13/pflag"
)


type AliceFlags struct {
	AliceConfigFile string `json:"alice_config_file"`
}

func NewAliceFlags() *AliceFlags {
	return &AliceFlags{}
}

func NewContainerRuntimeOptions()  {

}

func (a *AliceFlags) AddFlags(mainfs *pflag.FlagSet) {
	fs := pflag.NewFlagSet("", pflag.ExitOnError)
	defer func() {
		fs.VisitAll(func(f *pflag.Flag) {
			if len(f.Deprecated) > 0 {
				f.Hidden = false
			}
		})
		mainfs.AddFlagSet(fs)
	}()
	fs.StringVar(&a.AliceConfigFile, "config", a.AliceConfigFile, "te or relati")
}
