package cmd

import (
	"fmt"
	"os"

	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/queue"
	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/store"
	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/worker"
	"github.com/spf13/cobra"
)

var file string

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Runs the worker that retrieves dependency for a given package",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("worker called")
		if file == "" {
			os.Exit(1)
		}

		q, err := queue.NewMemoryQueueFromFile(file)
		if err != nil {
			panic(err)
		}

		s := store.NewMemoryStore()
		w := worker.NewWorker(q, s)
		if err := w.Start(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)

	workerCmd.Flags().StringVarP(&file, "file", "f", "", "File to read from")
}
