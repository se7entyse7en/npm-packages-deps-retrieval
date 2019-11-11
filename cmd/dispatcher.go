package cmd

import (
	"context"
	"fmt"

	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/dispatcher"
	"github.com/spf13/cobra"
)

var dispatcherCmd = &cobra.Command{
	Use:   "dispatcher",
	Short: "Runs the dispatcher that submits packages for which to retrieve dependencies",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dispatcher called")
		bootstrap, err := cmd.Flags().GetBool("bootstrap")
		if err != nil {
			panic(err)
		}

		topN, err := cmd.Flags().GetInt("topn")
		if err != nil {
			panic(err)
		}

		q, err := buildQueue(cmd, "file")
		if err != nil {
			panic(err)
		}

		ctx := context.Background()
		if err := q.Open(ctx); err != nil {
			panic(err)
		}

		defer q.Close(ctx)

		d := dispatcher.NewDispatcher(q, bootstrap, topN)
		if err := d.Start(ctx); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dispatcherCmd)

	dispatcherCmd.Flags().StringP("file", "f", "", "File to use as queue")
	dispatcherCmd.Flags().BoolP("bootstrap", "i", true, "Whether to do initial bootstrap")
	dispatcherCmd.Flags().IntP("topn", "t", 1000, "How many packages to bootstrap among most populars")
	dispatcherCmd.Flags().StringP("broker-uri", "b", "", "Broker uri of the RabbitMQ instance")
	dispatcherCmd.Flags().StringP("queue", "q", "", "Name of the RabbitMQ queue")
}
