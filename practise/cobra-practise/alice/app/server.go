package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"kube-learning/practise/cobra-practise/alice/app/flag"
	"kube-learning/practise/cobra-practise/alice/app/options"
)

const (
	// Kubelet component name
	componentAlice = "alice"
)

func NewAliceCommand() *cobra.Command {
	//使用 NewFlagSet 函数声明一个子命令， 然后为这个子命令定义一个专用的 flag
	cleanFlagSet := pflag.NewFlagSet(componentAlice, pflag.ContinueOnError)

	cleanFlagSet.SetNormalizeFunc(flag.WordSepNormalizeFunc)

	aliceFlags := options.NewAliceFlags()

	cmd := &cobra.Command{
		Use: componentAlice,
		Long: `alice test`,
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := cleanFlagSet.Parse(args); err != nil {
				return fmt.Errorf("failed to parse kubelet flag: %w", err)
			}
			fmt.Println(args)

			cmds := cleanFlagSet.Args()
			if len(cmds) > 0 {
				return fmt.Errorf("unknown command %+s", cmds[0])
			}
			return nil
		},
	}

	// keep cleanFlagSet separate, so Cobra doesn't pollute it with the global flags
	//添加 Flag
	aliceFlags.AddFlags(cleanFlagSet)

	const usageFmt = "Usage:\n  %s\n\nFlags:\n%s"
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine(), cleanFlagSet.FlagUsagesWrapped(2))
		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine(), cleanFlagSet.FlagUsagesWrapped(2))
	})

	return cmd
}