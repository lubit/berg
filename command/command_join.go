package command

import (
	"log"
	"strings"
	"time"
	"context"
	"os"

	"github.com/mitchellh/cli"
	"github.com/lubit/berg/rpc"
	"google.golang.org/grpc"
)

type CommandJoin struct {
	Ui cli.Ui
}

var _ cli.Command = &CommandJoin{}

func (c *CommandJoin) Run(args []string) int {

	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc conn err: %v", err)
	}
	defer conn.Close()

	cc := rpc.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := cc.StartJob(ctx, &rpc.JobRequest{Name: "berg"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting:%s", r.GetMessage())

	os.Exit(0)
	return 0
}

func (c *CommandJoin) Help() string {
	helpText := "berg join --"
	return strings.TrimSpace(helpText)
}
func (c *CommandJoin) Synopsis() string {
	return "Tell berg to join cluster"
}
