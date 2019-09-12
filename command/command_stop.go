package command

import (
	"strings"

	"github.com/lubit/berg/rpc"
	"github.com/mitchellh/cli"

	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

type CommandStop struct {
	Ui cli.Ui
}

func (c *CommandStop) Help() string {
	helpText := "berg stop --"
	return strings.TrimSpace(helpText)
}
func (c *CommandStop) Synopsis() string {
	return "Tell berg to stop"
}

func (c *CommandStop) Run(args []string) int {

	// new rpc server
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc conn err: %v", err)
	}
	defer conn.Close()

	cli := rpc.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := cli.StopJob(ctx, &rpc.JobRequest{Name: "berg"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting:%s", r.GetMessage())

	return 0
}
