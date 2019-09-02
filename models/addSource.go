package models

import (
	"context"
	"fmt"
	"sync"
	kafka "github.com/segmentio/kafka-go"
)


//定义一个数据接收器结构体
type DataSource struct {
	//同步方法集 任务队列不为0则阻塞程序继续运行
	wg           	sync.WaitGroup
	inChannel    	[]chan interface{}
	outChannel   	[]chan interface{}
	errorChannel 	[]chan interface{}
	exitChannel  	[]chan interface{}
	offsetposition		int64
	//将状态信息保存 传输给平台 展示数据状态、routine状态 {"data":"","routine":""}
	status       	interface{}
	//落地数据位置用来回滚
	barrier		 	interface{}
	//设定启动的go routine数量
	parallelism  	int
}

//增加数据接入配置 扩展给用户自定义
func (data *DataSource) AddSource(sourceconf interface{}) error{
	fmt.Printf("Input AddSource Value: %+v\n",sourceconf)
	return nil
}

//增加kafka接入配置
func (data *DataSource) AddKafka(kafkaconf interface{},wg *sync.WaitGroup)  (*kafka.Reader,error){

	//读取kafka配置信息
	brokers := []string{"1.0.0.0:9092",""}
	topic := ""
	groupid := ""
	partition := 0
	offsetnum := 99

	//消费者实例化 NewReader返回为指针
	consumer := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		GroupID:   groupid,
		Topic:     topic,
		Partition:	partition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	//自定义offset
	consumer.SetOffset(int64(offsetnum))

	//defer延迟执行 会在return后面执行
	defer consumer.Close()

	for {
		m, err := consumer.ReadMessage(context.Background())
		if err != nil {
			break
		}

		data.outChannel[0] <- m

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	fmt.Printf("Input AddSource Value: %+v\n",kafkaconf)
	return consumer,nil
}

//kafka指定offset
func (data *DataSource) SetOffset(in int64){
	data.offsetposition = in
	fmt.Printf("Input SetOffset Value: %+v\n",in)
}

//kafka运行的详细状态信息
func (data *DataSource) GetState() error{
	return nil
}

//设置多线程处理 routines数量
func (data *DataSource) SetParallel(in int) error{
	//设置运行时go多少次source()
	data.parallelism = in
	return nil
}

//设置channel大小
func (data *DataSource) SetChannelSize(in int) error{
	return nil
}

//传入Signal进行终止操作
func (data *DataSource) Terminate(){

}

//前面信息初始化后进行数据源加载启动
func (data *DataSource) Execute() error{

	//启动routine运行

	//启动kafka消费

	//s




	return nil
}