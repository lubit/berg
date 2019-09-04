package models

import "fmt"

//DataSink结构体
type DataSink struct{

}

//初始化DataSink
func NewDataSink() *DataSink{
	fmt.Println("New DataSink")
	return &DataSink{}
}