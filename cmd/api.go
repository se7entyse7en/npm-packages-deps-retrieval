package cmd

import (
	"context"
	"fmt"
	"net"

	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Runs the api that answers to requests for packages deps",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("api called")
		s, err := buildStore(cmd, "file")
		if err != nil {
			panic(err)
		}

		ctx := context.Background()
		if err := s.Open(ctx); err != nil {
			panic(err)
		}

		defer s.Close(ctx)

		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			panic(err)
		}

		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			panic(err)
		}

		grpcServer := grpc.NewServer()
		server := api.NewApiServer(s)
		api.RegisterDependenciesServiceServer(grpcServer, server)
		reflection.Register(grpcServer)
		fmt.Printf("listening on port %d\n", port)
		grpcServer.Serve(lis)
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	apiCmd.Flags().StringP("file", "f", "", "File to use as storage")
	apiCmd.Flags().IntP("port", "p", 8080, "Port to listen to")
	apiCmd.Flags().StringP("db-uri", "u", "", "DB uri of the MongoDB instance to use as storage")
	apiCmd.Flags().StringP("db", "d", "", "Database name to write output to use as storage")
	apiCmd.Flags().StringP("coll", "c", "", "Collection name to write output to use as storage")
}
