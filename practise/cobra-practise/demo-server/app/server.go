package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

func NewDemoCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:  "demo-server",
		Long: `The demo server controller is a daemon than embeds the core control loops shipped with demo.`,
		Run: func(cmd *cobra.Command, args []string) {
			http.HandleFunc("/hello", hello)
			http.HandleFunc("/headers", headers)

			http.ListenAndServe(":8090", nil)
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

	return cmd
}


func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
