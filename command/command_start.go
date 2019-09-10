package command

import (
	"fmt"
	"strings"

	"github.com/hashicorp/memberlist"
	"github.com/hashicorp/serf/serf"
	"github.com/lubit/berg/rpc"
	"github.com/lubit/berg/web"
	"github.com/mitchellh/cli"
)

type CommandStart struct {
	Ui cli.Ui
}

func (c *CommandStart) Run(args []string) int {

	// new gossip cluster
	go NewSerf()
	// new rpc server
	go rpc.NewRPCServer()
	// web server
	go web.NewWebServer()

	return 0
}
func (c *CommandStart) Help() string {
	helpText := "berg start --"
	return strings.TrimSpace(helpText)
}
func (c *CommandStart) Synopsis() string {
	return "Tell berg to start"
}

func NewSerf() {
	//config
	//start && listen && join pri LAN
	//
	serfConfig := serf.DefaultConfig()
	serfConfig.MemberlistConfig = memberlist.DefaultLocalConfig()
	serfConfig.MemberlistConfig.BindAddr = "0.0.0.0"
	serfConfig.MemberlistConfig.BindPort = 8002

	fmt.Printf("agent conf: %+v \n", serfConfig.MemberlistConfig)
	//conf.Init()
	berf, err := serf.Create(serfConfig)
	if err != nil {
		fmt.Println("failed to start lan serf: ", err)
	}

	//berf.Shutdown()

	shutdownCH := berf.ShutdownCh()
	for {
		select {
		case <-shutdownCH:
			fmt.Println("shutdown")
			return

		}
	}
}
