package models

import "fmt"

//DataProcessor结构体
type DataProcessor struct{

}

//初始化DataSource
func NewDataProcessor() *DataProcessor{
	fmt.Println("New DataProcessor")
	return &DataProcessor{}
}