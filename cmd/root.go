package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "dd-scraper",
	Short: "Poll data from datadog api",
	Long:  "Poll data from datadog api",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Basic commands TODO")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
