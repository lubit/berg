package main

import (
	"fmt"
	"time"

	pb "github.com/lubit/berg/rpc"
	"github.com/shirou/gopsutil/mem"
)

/*
	1. 起停服务 带回调函数
	2. 加载配置
	3. web接口/rpc接口
	4. node
*/
func main() {
	// config
	sysinfo()
	go pb.NewRPCServer()
	// MainService()
	// web
	NewDeath(final).Wait()
}

//TODO 回收函数
func final(dch *chan struct{}) {
	defer close(*dch)
	fmt.Println("berg finalize ")
	time.Sleep(5 * time.Second)
	return
}

func sysinfo() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
}
