package cmd

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/api"
	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/store"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var file string
var port int

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Runs the api that answers to requests for packages deps",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("api called")
		if file == "" {
			os.Exit(1)
		}

		s := store.NewFileStore(file, false)
		ctx := context.Background()
		if err := s.Open(ctx); err != nil {
			panic(err)
		}

		defer s.Close(ctx)

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

	apiCmd.Flags().StringVarP(&file, "file", "f", "", "File to use as storage")
	apiCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to listen to")
}
