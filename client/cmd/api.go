package cmd

import (
	"fmt"
	"context"
	"time"

	pb "github.com/zulcss/edged/proto"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var apiCommand = &cobra.Command{
	Use: 	"api",
	Short:	"Query Edged API worker",
}

var apiShow = &cobra.Command{
	Use:	"ping",
	Short:	"Show Edged API information",
	Run: func(cmd *cobra.Command, args[] string) {
		conn, err := grpc.Dial(fmt.Sprintf("%s:8000", Server),
				grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Println("Failed to connect to: %v", err)
		}
		defer conn.Close()
		c := pb.NewEdgeClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.Ping(ctx,  &pb.VersionRequest{Server: fmt.Sprintf("%s", Server)})
		if err != nil {
			fmt.Println("%s is down", Server)
		}
		if r.Message != "" {
			fmt.Printf("%s is up\n", Server)
		}
	},
}

func init() {
	rootCmd.AddCommand(apiCommand)

	apiCommand.AddCommand(apiShow)

}
