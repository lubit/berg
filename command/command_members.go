package command

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/lubit/berg/rpc"
	"github.com/mitchellh/cli"
	"google.golang.org/grpc"
)

type CommandMembers struct {
	Ui cli.Ui
}

func (c *CommandMembers) Help() string {
	helpText := "berg stop --"
	return strings.TrimSpace(helpText)
}
func (c *CommandMembers) Synopsis() string {
	return "Tell berg to stop"
}

func (c *CommandMembers) Run(args []string) int {

	// new rpc server
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc conn err: %v", err)
	}
	defer conn.Close()

	cli := rpc.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := cli.Members(ctx, &rpc.JobRequest{Name: "members"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("members:%s", r.GetMessage())

	return 0
}
