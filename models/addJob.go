package models

import (
	"sync"
)

//Application结构体作为核心模块 提供附加各种状态信息 以及 定制动能
/*
channel用来控制Channel的类型、数量以及大小
*/
type Application struct{
	/*
	流程控制模块
	 */
	//Channel&Routine结构体
	AppChannels 					*Channel
	//同步方法集 任务队列不为0则阻塞程序继续运行
	AppWaitGroup           			sync.WaitGroup
	//DataSource结构体
	AppSource 						*DataSource
	//DataProcessor结构体
	AppProcessor					*DataProcessor
	//DataSink结构体
	AppSink 						*DataSink

	/*
	运行状态信息模块
	 */
	//任务运行类型
	AppType							string
	//任务运行状态
	AppStatus						string
}

//初始化Application
func NewJob() *Application {
	var app Application
	app.AppChannels = NewChannel()
	app.AppSource = NewDataSource()
	app.AppProcessor = NewDataProcessor()
	app.AppSink = NewDataSink()
	return &app
}

//Application的资源、运行逻辑设置&准备
func (app *Application) Prepare(){

}

//Application的启动
func (app *Application) Execute(){

}

//Application的终止
func (app *Application) Terminate(){
	//关闭所有Channel
	app.AppChannels.Terminate()
}

/*
一个application需要哪些必要模块来启动
1 - 启动一个application实例
2 - 启动一个kafka数据消费实例
3 - 启动一个channel用来数据传输
4 - 将kafka消息导向channel中
5 - 启动一个Sink读取channel中数据并写入到目标数据库
*/