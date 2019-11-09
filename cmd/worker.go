package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/queue"
	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/store"
	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/worker"
	"github.com/spf13/cobra"
)

var inFile string
var outFile string

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Runs the worker that retrieves dependency for a given package",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("worker called")
		if inFile == "" {
			os.Exit(1)
		}

		q, err := queue.NewMemoryQueueFromFile(inFile)
		if err != nil {
			panic(err)
		}

		s := store.NewFileStore(outFile, true)
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

	workerCmd.Flags().StringVarP(&inFile, "in-file", "i", "", "File to read input from")
	workerCmd.Flags().StringVarP(&outFile, "out-file", "o", "", "File to write output t")
}
