package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Runs the api that answers to requests for packages deps",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("api called")
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
