package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	"kube-learning/practise/cobra-practise/demo-server/app/options"
	"os"
)

var Source string
func NewDemoCommand() *cobra.Command {
	o, err := options.NewOptions()
	if err != nil {
		klog.Fatalf("unable to initialize command options: %v", err)
	}

	cmd := &cobra.Command{
		Use:  "demo-server",
		Long: `The demo server controller is a daemon than embeds the core control loops shipped with demo.`,
		Run: func(cmd *cobra.Command, args []string) {
			fs := cmd.Flags()
			o.Fs = fs
			if err = o.Complete(); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			if err = o.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	fs := cmd.Flags()
	namedFlagSets := o.Flags()
	for _, f := range namedFlagSets.FlagSets {
		fs.AddFlagSet(f)
	}

	return cmd
}