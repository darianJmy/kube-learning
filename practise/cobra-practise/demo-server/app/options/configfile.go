package options

import (
	"github.com/spf13/pflag"
	"k8s.io/component-base/cli/globalflag"
	"os"

	"gopkg.in/yaml.v2"

	"kube-learning/practise/cobra-practise/demo-server/app/config"
)

func loadConfigFromFile(file string) (*config.Config, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return loadConfig(data)
}

func loadConfig(data []byte) (*config.Config, error) {
	var c config.Config
	if err := yaml.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

func AddCustomGlobalFlags(fs *pflag.FlagSet) {
	// Lookup flags in global flag set and re-register the values with our flagset.

	// Adds flags from k8s.io/kubernetes/pkg/cloudprovider/providers.
	registerLegacyGlobalFlags(fs)

	// Adds flags from k8s.io/apiserver/pkg/admission.
	globalflag.Register(fs, "default-not-ready-toleration-seconds")
	globalflag.Register(fs, "default-unreachable-toleration-seconds")
}

func registerLegacyGlobalFlags(fs *pflag.FlagSet) {
	globalflag.Register(fs, "cloud-provider-gce-lb-src-cidrs")
	globalflag.Register(fs, "cloud-provider-gce-l7lb-src-cidrs")
	fs.MarkDeprecated("cloud-provider-gce-lb-src-cidrs", "This flag will be removed once the GCE Cloud Provider is removed from kube-apiserver")
}
