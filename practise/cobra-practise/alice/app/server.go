package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
	"kube-learning/practise/cobra-practise/alice/app/flag"
	"kube-learning/practise/cobra-practise/alice/app/options"
	"os"
	"time"
)

const (
	// Kubelet component name
	componentAlice = "alice"
)

func NewAliceCommand() *cobra.Command {
	//使用 NewFlagSet 函数声明一个子命令， 然后为这个子命令定义一个专用的 flag
	cleanFlagSet := pflag.NewFlagSet(componentAlice, pflag.ContinueOnError)
	//设置标准化参数名
	cleanFlagSet.SetNormalizeFunc(flag.WordSepNormalizeFunc)
	//初始化aliceFlags
	aliceFlags := options.NewAliceFlags()

	cmd := &cobra.Command{
		Use:                componentAlice,
		Long:               `alice test`,
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cleanFlagSet.Parse(args); err != nil {
				klog.ErrorS(err, "Failed to parse alice flag")
				//print 参数
				cmd.Usage()
				os.Exit(1)
			}

			cmds := cleanFlagSet.Args()
			if len(cmds) > 0 {
				klog.ErrorS(nil, "Unknown command", "command", cmds[0])
				cmd.Usage()
				os.Exit(1)
			}

			help, err := cleanFlagSet.GetBool("help")
			if err != nil {
				klog.InfoS(`"help" flag is non-bool, programmer error, please correct`)

			}
			if help {
				//print 参数
				cmd.Help()
				os.Exit(1)
			}

			if err := options.ValidateAliceFlags(aliceFlags); err != nil {
				klog.ErrorS(err, "Failed to validate alice flags")
				os.Exit(1)
			}
			i := 0
			stopCh := make(chan struct{})
			go wait.Until(func() {
				fmt.Printf("----%d---\n", i)
				i++
			}, time.Second*3, stopCh)

			for {
				time.Sleep(3 * time.Second)
				fmt.Printf("ok")

			}
			return nil
		},
	}

	//添加 Flag
	aliceFlags.AddFlags(cleanFlagSet)
	//添加 Help Flag
	cleanFlagSet.BoolP("help", "h", false, fmt.Sprintf("help for %s", cmd.Name()))

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
