package options

import (
	"github.com/spf13/pflag"
)


type AliceFlags struct {
	AliceConfigFile string
	Help bool
}

func NewAliceFlags() *AliceFlags {
	return &AliceFlags{
	}
}

func ValidateAliceFlags(f *AliceFlags) error {

	return nil
}

func NewContainerRuntimeOptions()  {

}

func (a *AliceFlags) AddFlags(mainfs *pflag.FlagSet) {
	fs := pflag.NewFlagSet("", pflag.ExitOnError)
	defer func() {
		//遍历所有定义的参数
		fs.VisitAll(func(f *pflag.Flag) {
			if len(f.Deprecated) > 0 {
				f.Hidden = false
			}
		})
		mainfs.AddFlagSet(fs)
	}()
	fs.StringVar(&a.AliceConfigFile, "config-file", a.AliceConfigFile, "te or relati")
}