package docker

import (
	"github.com/spf13/cobra"
)

type InitDockerOptions struct {
	Name       string   `json:"name"`
	ValuesSets []string `json:"value"`
}

func NewInitDockerOptions() *InitDockerOptions {
	return &InitDockerOptions{}
}

func Docker(initOptions *InitDockerOptions) *cobra.Command {
	if initOptions == nil {
		initOptions = NewInitDockerOptions()
	}

	cmd := &cobra.Command{
		Use:   "docker",
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
	}
	return cmd
}

func run(initOptions *InitDockerOptions) error {
	err := checkDocker()
	return nil
}

func checkDocker() {

}
