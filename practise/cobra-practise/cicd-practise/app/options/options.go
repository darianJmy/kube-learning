package options

import "github.com/spf13/pflag"

type CicdFlags struct {
}

func NewCicdFlags() *CicdFlags {
	return &CicdFlags{}
}

func (a *CicdFlags) AddFlags(mainfs *pflag.FlagSet) {
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
}
