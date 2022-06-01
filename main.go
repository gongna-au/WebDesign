package main

import (
	"fmt"
	//config "github.com/LeetCode/LibBook/DB/config"
	model "github.com/WebDesign/model"
	"time"
)

func main() {
	//配置文件初始化
	//config.ConfigInit()
	//数据的初始化
	DB1 := model.GetDBInstance()
	DB2 := model.GetDBInstance()
	if DB1 == DB2 {
		fmt.Println("equal")
	}

	fmt.Println("main wait ")
	time.Sleep(time.Second * 30)
}
