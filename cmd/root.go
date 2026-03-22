package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "weclaw",
	Short: "WeChat AI agent bridge",
	Long:  "weclaw bridges WeChat messages to AI agents via the iLink API.",
	RunE:  runStart, // default command is start
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
