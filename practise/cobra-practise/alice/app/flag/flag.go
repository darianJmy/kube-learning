package flag

import (
	"github.com/spf13/pflag"
	"strings"
)

// WordSepNormalizeFunc changes all flags that contain "_" separators
func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "-") {
		return pflag.NormalizedName(strings.Replace(name, "-", "_", -1))
	}
	return pflag.NormalizedName(name)
}