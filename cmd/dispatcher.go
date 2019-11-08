package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dispatcherCmd = &cobra.Command{
	Use:   "dispatcher",
	Short: "Runs the dispatcher that submits packages for which to retrieve dependencies",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dispatcher called")
	},
}

func init() {
	rootCmd.AddCommand(dispatcherCmd)
}
