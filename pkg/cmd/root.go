package cmd

import (
	"fmt"
	"github.com/marcnuri-work/go-ssr/pkg/web"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "go-ssr",
	Short: "Server Side Rendering (SSR) for Go applications demo",
	Run: func(cmd *cobra.Command, args []string) {
		web.StartServer()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error executing command: %s", err)
		os.Exit(1)
	}
}
