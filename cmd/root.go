package cmd

import (
	"fmt"
	"os"

	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/queue"
	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/store"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "npm-packages-deps-retrieval",
	Short: "A CLI to spin up the components for running the npm-packages-deps-retrieval",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func buildStore(cmd *cobra.Command, fileParamName string) (store.Store, error) {
	dbURI, err := cmd.Flags().GetString("db-uri")
	if err != nil {
		return nil, err
	}

	var s store.Store
	if dbURI != "" {
		db, err := cmd.Flags().GetString("db")
		if err != nil {
			return nil, err
		}

		coll, err := cmd.Flags().GetString("coll")
		if err != nil {
			return nil, err
		}

		s, err = store.NewMongoStore(dbURI, db, coll)
		if err != nil {
			return nil, err
		}
	} else {
		file, err := cmd.Flags().GetString(fileParamName)
		if err != nil {
			return nil, err
		}

		if file == "" {
			return nil, fmt.Errorf("empty file name")
		}

		s = store.NewFileStore(file, true)
	}

	return s, nil
}

func buildQueue(cmd *cobra.Command, fileParamName string) (queue.Queue, error) {
	brokerURI, err := cmd.Flags().GetString("broker-uri")
	if err != nil {
		return nil, err
	}

	var q queue.Queue
	if brokerURI != "" {
		qName, err := cmd.Flags().GetString("queue")
		if err != nil {
			return nil, err
		}

		q, err = queue.NewRabbitMQQueue(brokerURI, qName)
		if err != nil {
			return nil, err
		}
	} else {
		file, err := cmd.Flags().GetString(fileParamName)
		if err != nil {
			return nil, err
		}

		if file == "" {
			return nil, fmt.Errorf("empty file name")
		}

		q = queue.NewFileQueue(file, true)
	}

	return q, nil
}
