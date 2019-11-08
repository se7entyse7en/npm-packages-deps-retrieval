package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Runs the worker that retrieves dependency for a given package",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("worker called")
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
