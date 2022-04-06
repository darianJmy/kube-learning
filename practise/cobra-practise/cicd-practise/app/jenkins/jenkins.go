package jenkins

import (
	"fmt"
	"github.com/spf13/cobra"
)

func Jenkins() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "jenkins",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello2")
		},
	}
	return cmd
}
