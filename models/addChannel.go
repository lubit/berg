package models

//Channel的handler用来定义传输过程中需要多少channel以及数据从哪一个channel交互
type Channel struct{
	ChannelMap map[string][]chan interface{}
}

//注册Channel的方法 输入Channel名字以及Channel大小
func (channel *Channel) AddChannel(channelname string, channelnumber int64){
	channel.ChannelMap[channelname] = make([]chan interface{},channelnumber)
}


