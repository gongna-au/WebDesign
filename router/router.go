package router

import (
	"net/http"
	"strings"

	"github.com/WebDesign/handler/login"
	"github.com/WebDesign/handler/signup"
	"github.com/WebDesign/handler/todo"
	"github.com/WebDesign/handler/user"

	"github.com/WebDesign/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {

	// 注册全局中间件
	registerGlobalMiddleWare(router)

	//  注册 API 路由
	RegisterAPIRoutes(router)

	//  配置 404 路由
	setup404Handler(router)
}

//注册全局中间件
func registerGlobalMiddleWare(g *gin.Engine) {
	g.Use(middlewares.Logger())

}

//  注册 API 路由
func RegisterAPIRoutes(g *gin.Engine) {

	g1 := g.Group("/api/v1/todos")
	{
		g1.POST("/", todo.CreateTodo)
		g1.GET("/", todo.FetchAllTodo)
		g1.GET("/:id", todo.FetchSingleTodo)
		g1.PUT("/:id", todo.UpdateTodo)
		g1.DELETE("/:id", todo.DeleteTodo)
	}

	g2 := g.Group("/api/v1/auth")
	{

		//使用手机号和密码注册
		g2.POST("/signup/usingphone", signup.SignupUsingPhone)
		//手机号是否已注册
		g2.POST("/signup/phone/exist", signup.IsPhoneExist)

	}

	g3 := g.Group("/api/v1/login")
	{

		g3.POST("/usingphone", login.LoginByPhone)

	}

	g4 := g.Group("/api/v1")
	{

		//获取当前用户
		g4.GET("/user", user.GetCurrentUser)
		//用户列表
		g4.GET("/users", user.GetUsers)
		//修改个人资料
		g4.PUT("/users", user.UpdateProfile)

	}

}

//  配置 404 路由
func setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})

}
