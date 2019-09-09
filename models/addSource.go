package models

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"golang.org/x/net/context"
	"sync"
)

//定义一个数据接收器结构体
type DataSource struct {
	//同步方法集 任务队列不为0则阻塞程序继续运行
	wg           			sync.WaitGroup
	//自定义offset位置
	offsetposition			int64
	//将状态信息保存 传输给平台 展示数据状态、routine状态 {"data":"","routine":""}
	status       			interface{}
	//落地数据位置用来回滚
	barrier		 			interface{}
	//设定启动的go routine数量
	parallelism  			int
	//设定传出channel
	outputchannel			[]chan interface{}

	consumer				*kafka.Reader
}

//初始化DataSource
func NewDataSource() *DataSource{
	fmt.Println("New DataSource")
	return &DataSource{}
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
	data.consumer = kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		GroupID:   groupid,
		Topic:     topic,
		Partition:	partition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	//自定义offset
	data.consumer.SetOffset(int64(offsetnum))

	fmt.Printf("Input AddSource Value: %+v\n",kafkaconf)
	return data.consumer,nil
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

func (data *DataSource) Run(inchan chan interface{}, outchan chan interface{}){

	defer data.consumer.Close()

	//利用context来接收退出信号
	ctx,cancel := context.WithCancel(context.Background())

	go func(){
		defer cancel()
		<- inchan
	}()

	for {
		m, err := data.consumer.ReadMessage(ctx)
			if err != nil && err == ctx.Err(){
				break
			}

			outchan <- m

			fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

	}
}


//前面信息初始化后进行数据源加载启动
func (data *DataSource) Execute(chns *Channel) error{

	//

	//启动routine运行

	//启动kafka消费

	//写入outChannel

	for i := 0; i < len(chns.ChannelQueue[1]); i++{
		go data.Run(chns.ChannelQueue[0][i],chns.ChannelQueue[1][i])
	}


	return nil
}