package main

import (
	"fmt"
	"time"

	"github.com/lubit/berg/rpc"
	"github.com/lubit/berg/web"
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
	// config

	go rpc.NewRPCServer()
	go web.NewWebServer()
	// MainService()
	// web
	//go sysinfo()
	//trace.Start(os.Stderr)
	//defer trace.Stop()

	NewDeath(finalize).Wait()
}

//TODO 回收函数
func finalize(dch *chan struct{}) {
	defer close(*dch)
	fmt.Println("berg finalize ")
	time.Sleep(1 * time.Second)
	return
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
