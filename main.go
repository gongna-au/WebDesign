package main

import (
	//"fmt"
	//_ "github.com/WebDesign/handler/login"
	//_ "github.com/WebDesign/handler/signup"
	_ "github.com/WebDesign/handler/todo"
	//_ "github.com/WebDesign/handler/user"

	_ "github.com/WebDesign/config"
	_ "github.com/WebDesign/log"

	_ "github.com/WebDesign/database"
	"github.com/WebDesign/router"

	"github.com/gin-gonic/gin"
	//"github.com/WebDesign/model"
	//"go.uber.org/zap"
	//"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin"
	//"time"
)

func main() {

	g := gin.New()
	//路由初始化
	router.SetupRoute(g)
	g.Run()

}
