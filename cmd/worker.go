package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/queue"
	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/worker"
	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Runs the worker that retrieves dependency for a given package",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("worker called")
		inFile, err := cmd.Flags().GetString("in-file")
		if err != nil {
			panic(err)
		}

		if inFile == "" {
			os.Exit(1)
		}

		q, err := queue.NewMemoryQueueFromFile(inFile)
		if err != nil {
			panic(err)
		}

		s, err := buildStore(cmd, "out-file")
		if err != nil {
			panic(err)
		}

		ctx := context.Background()
		if err := s.Open(ctx); err != nil {
			panic(err)
		}

		defer s.Close(ctx)

		w := worker.NewWorker(q, s)
		if err := w.Start(ctx); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)

	workerCmd.Flags().StringP("in-file", "i", "", "File to read input from")
	workerCmd.Flags().StringP("out-file", "o", "", "File to write output to")
	workerCmd.Flags().StringP("db-uri", "u", "", "DB uri of the MongoDB instance to write output to")
	workerCmd.Flags().StringP("db", "d", "", "Database name to write output to")
	workerCmd.Flags().StringP("coll", "c", "", "Collection name to write output to")
}
