package main

import (
	"go-svc-tpl/cmd"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	store := cookie.NewStore([]byte("secret"))
	// 创建 gin 引擎对象
	r := gin.Default()

	// 使用 sessions 中间件
	r.Use(sessions.Sessions("session", store))

	// 注册路由和控制器
	//待补充

	// 启动服务器
	r.Run(":8080")
	/////////////////////////////
	cmd.Execute()
}
