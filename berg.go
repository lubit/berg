package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lubit/berg/command"
	"github.com/mitchellh/cli"
)

/*
	1. 起停服务
		1.1 Finalize带回调函数
		1.2 Job | Group 起停
	2. 加载配置
		2.1 所有可运行任务的配置化
	3. web接口/rpc接口
	4. 资源隔离 slot?
	5. 分布式任务管理
		5.1 任务形式：plugin
		5.2 任务交互：http2 file + conf
*/
func main() {

	NewClient()
	// config
	// rpc server

	// gossip cluster

	// MainService()
	// web
	//go sysinfo()
	//trace.Start(os.Stderr)
	//defer trace.Stop()
}

func NewClient() {

	ui := &cli.BasicUi{Writer: os.Stdout}

	c := cli.NewCLI("berg", "0.0.1")
	c.Commands = map[string]cli.CommandFactory{
		"start": func() (cli.Command, error) { return &command.CommandStart{Ui: ui}, nil },
		"stop":  func() (cli.Command, error) { return &command.CommandStop{Ui: ui}, nil },
		"join":  func() (cli.Command, error) { return &command.CommandJoin{Ui: ui}, nil },
	}
	c.Args = os.Args[1:]
	_, err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("berg command finished")
}

func sysinfo() {
	/*
		v, _ := mem.VirtualMemory()
		// almost every return value is a struct
		fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
		// convert to JSON. String() is also implemented
		fmt.Println(v)
	*/
	/*
		mux := http.NewServeMux()
		mux.HandleFunc("/profile", pprof.Profile)
		log.Fatal(http.ListenAndServe(":7777", mux))
	*/
	/*
		trace.Start(os.Stderr)
		defer trace.Stop()
	*/
}
