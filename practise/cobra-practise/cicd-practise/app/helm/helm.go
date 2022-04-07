package helm

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"log"
	"os/exec"
)

const (
	cmdHelm            string = "helm"
	opInstall          string = "install"
	helmName           string = "harbor"
	helmChartReference string = "bitnami/harbor"
)

type InitHelmOptions struct {
	Name           string   `json:"name"`
	ChartReference string   `json:"chart_reference"`
	ValuesSets     []string `json:"value"`
}

func NewInitHelmOptions() *InitHelmOptions {
	return &InitHelmOptions{}
}

func Helm(initOptions *InitHelmOptions) *cobra.Command {
	if initOptions == nil {
		initOptions = NewInitHelmOptions()
	}

	cmd := &cobra.Command{
		Use:   "helm",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(initOptions)
		},
		Args: cobra.NoArgs,
	}

	AddInitHelmFlags(cmd.Flags(), initOptions)

	return cmd
}

func AddInitHelmFlags(flagSet *pflag.FlagSet, init *InitHelmOptions) {
	flagSet.StringVar(
		&init.Name, "name", init.Name,
		"helm name",
	)
	flagSet.StringVar(
		&init.ChartReference, "chartReference", init.ChartReference,
		`helm chartreference`,
	)
	flagSet.StringArrayVar(
		&init.ValuesSets, "set", init.ValuesSets,
		"helm options",
	)

}

func run(init *InitHelmOptions) error {
	err := checkHelm()
	if err != nil {
		return fmt.Errorf("helm is not found")
	}
	if init.Name == "" && init.ChartReference == "" {
		return fmt.Errorf("name or chartReference is not found")
	}
	args := loadConfigFile(init)

	out, err := exec.Command(cmdHelm, args...).CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return err
	}
	fmt.Printf("combined out:\n%s\n", string(out))
	return nil
}

func loadConfigFile(cfg *InitHelmOptions) []string {
	args := []string{opInstall, cfg.Name, cfg.ChartReference}
	if &cfg.ValuesSets == nil {
		var ValuesSet map[string]string
		ValuesSet = make(map[string]string)
		ValuesSet["service.nodePorts.http"] = "30001"
		ValuesSet["global.storageClass"] = "managed-nfs-storage"
		ValuesSet["service.tls.enabled"] = "false"
		ValuesSet["externalURL"] = "http://10.10.33.34:30001"

		for k, v := range ValuesSet {
			args = append(args, []string{"--set", fmt.Sprintf("%s=%s", k, v)}...)
		}
	} else {
		for _, v := range cfg.ValuesSets {
			args = append(args, []string{"--set", v}...)
		}
	}

	return args
}

func checkHelm() error {
	err := exec.Command("which", cmdHelm).Run()
	if err != nil {
		return fmt.Errorf("helm is not found")
	}

	err = exec.Command(cmdHelm, "list").Run()
	if err != nil {
		return fmt.Errorf("helm is not authenticate")
	}
	return nil
}
