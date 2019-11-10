package cmd

import (
	"fmt"
	"os"

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
		outFile, err := cmd.Flags().GetString(fileParamName)
		if err != nil {
			return nil, err
		}

		s = store.NewFileStore(outFile, true)
	}

	return s, nil
}
