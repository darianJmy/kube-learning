package harbor

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"io"
	"log"
	"os/exec"
)

const (
	cmdHelm   string = "helm"
	opInstall string = "install"
)

type InitHarborOptions struct {
	Name           string   `json:"name"`
	ChartReference string   `json:"chart_reference"`
	ValuesSets     []string `json:"value"`
}

// newInitOptions returns a struct ready for being used for creating cmd init flags.
func NewInitHarborOptions() *InitHarborOptions {
	return &InitHarborOptions{}
}

func Harbor(out io.Writer, initOptions *InitHarborOptions) *cobra.Command {
	if initOptions == nil {
		initOptions = NewInitHarborOptions()
	}

	cmd := &cobra.Command{
		Use:   "harbor",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(initOptions)
		},
		Args: cobra.NoArgs,
	}
	AddInitHarborFlags(cmd.Flags(), initOptions)

	return cmd
}

func AddInitHarborFlags(flagSet *pflag.FlagSet, init *InitHarborOptions) {
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
		"helm set",
	)

}

func run(init *InitHarborOptions) error {
	args := []string{opInstall, init.Name, init.ChartReference}
	if len(init.ValuesSets) != 0 {
		for _, v := range init.ValuesSets {
			args = append(args, []string{"--set", v}...)
		}
	}
	fmt.Println(args)
	out, err := exec.Command(cmdHelm, args...).CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return err
	}
	fmt.Printf("combined out:\n%s\n", string(out))
	return nil
}
