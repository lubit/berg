package models

import (
	"fmt"
	"sync"
)

//Channel的handler用来定义传输过程中需要配置多少Channel以及数据从哪一个Channel交互
//需要多少Channel 就需要多少Rountines; Channel&Routine一一对应
type Channel struct{
	//利用Channel Queue来顺序储存要使用的Channel队列
	ChannelQueue [][]chan interface{}
	//
	sync.RWMutex
	//ChannelMap默认必须有的是default的channel以及routine数量用来做单纯的数据接入以及写入 从配置中获得
	DefaultChannelNumber				int
	DefaultChannelSize					int
}

//初始化Channel
func NewChannel() *Channel{
	var channel Channel

	//加载Close Signal Channel位置为0
	channel.ChannelQueue = append(channel.ChannelQueue, make([]chan interface{}, channel.DefaultChannelNumber))
	for i := 0; i < channel.DefaultChannelNumber; i++ {
		channel.ChannelQueue[0][i] = make(chan interface{},channel.DefaultChannelSize)
	}

	//加载Source Channel位置为1
	channel.ChannelQueue = append(channel.ChannelQueue, make([]chan interface{}, channel.DefaultChannelNumber))
	for i := 0; i < channel.DefaultChannelNumber; i++ {
		channel.ChannelQueue[1][i] = make(chan interface{},channel.DefaultChannelSize)
	}

	fmt.Println("New Channel Initiated")
	return &channel
}

//注册Channel的方法 输入Channel数量以及Channel大小
/*
demo样例
ChannelMap里面包含三个主Channel
1 - ChannelMap[0][channelsize]
	--
2 - ChannelMap[1][channelsize]
3 - ChannelMap[2][channelsize]
 */
func (channel *Channel) AddChannel(channelnumber int64 , channelsize int64){
	channel.Lock()
	defer channel.Unlock()
	channel.ChannelQueue = append(channel.ChannelQueue, make([]chan interface{}, channelnumber))

	channelpos := len(channel.ChannelQueue) - 1

	for i:=0;i<len(channel.ChannelQueue[channelpos]);i++{
		channel.ChannelQueue[channelpos][i] = make(chan interface{}, channelsize)
	}

	fmt.Printf("AddChannel() Channel Queue Position:%s, Channel Number:%d, Channel Size:%d",
		channelpos, channelnumber, channelsize)
}

//进行各Channel终止操作
func (channel *Channel) Terminate(){
	channel.Lock()
	defer channel.Unlock()

	//顺序遍历各channel并依次关闭
	for i := 0; i < len(channel.ChannelQueue); i++{
		for pos := 0; pos < len(channel.ChannelQueue[i]); pos++ {
			close(channel.ChannelQueue[i][pos])
		}
	}

}