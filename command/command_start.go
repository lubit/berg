package command

import (
	"fmt"
	"strings"
	"time"

	"github.com/mitchellh/cli"
)

type CommandStart struct {
	Ui cli.Ui
}

func (c *CommandStart) Run(args []string) int {

	/*

	 */
	agent := &agent{}
	agent.Start()
	//NewDeath(finalize).Wait()

	return 0
}

//TODO 回收函数
func finalize(dch *chan struct{}) {
	defer close(*dch)
	fmt.Println("berg finalize ")
	time.Sleep(1 * time.Second)
	return
}

func (c *CommandStart) Help() string {
	helpText := "berg start --"
	return strings.TrimSpace(helpText)
}
func (c *CommandStart) Synopsis() string {
	return "Tell berg to start"
}

/*
func NewSerf() {
	//config
	//start && listen && join pri LAN
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
*/
