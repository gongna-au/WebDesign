package main

import (
	"fmt"
	"github.com/WebDesign/config"
	//"github.com/WebDesign/log"
	"github.com/WebDesign/model"
	//"go.uber.org/zap"

	"time"
)

func main() {

	DB1 := model.GetDBInstance()
	DB2 := model.GetDBInstance()
	if DB1 == DB2 {
		fmt.Println("equal")
	}
	c := config.GetConfigInfo()
	fmt.Println(c.MySQL)
	fmt.Println("main wait ")
	//log.Info("ferjwfgnerk", zap.String("get", "tettttt"))
	time.Sleep(time.Second * 30)
}
