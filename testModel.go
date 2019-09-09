package main

import (
	"context"
	"fmt"
	//"github.com/lubit/berg/models"
)

func TestModel() {
	//application := models.NewJob()
	////application.AppChannels.AddChannel()
	//
	//application.AppSource.Execute(application.AppChannels)

	var tmpslice []int
	tmpslice = append(tmpslice, 1)
	tmpslice = append(tmpslice, 2)

	fmt.Println(tmpslice[0], tmpslice[1])

	//fmt.Printf("New Application Initiated:%v",application)

	a()

	inchan := make(chan int, 1)

	inchan <- 1

	ctx, cancel := context.WithCancel(context.Background())

	//go func(){
	//	defer cancel()
	//	<- inchan
	//	fmt.Println("inside")
	//}()

	fmt.Println(ctx, cancel)
	//fearless-ry added for name test
	//context
}

func a() {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	fmt.Println("value i:", i)
}
