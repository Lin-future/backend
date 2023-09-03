package route

import (
	"go-svc-tpl/api/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Resp{
		Code: http.StatusOK,
		Msg:  "success",
		Data: "pong~",
	})
}

func SetupRouter(r *gin.RouterGroup) {
	r.GET("/ping", Ping)
	r.POST("/register", Register)
	r.POST("/login", Login)

	setupUserController(r)
}

func Register(ctx *gin.Context) {
	// 注册
	// ctx.JSON(http.StatusOK, dto.Resp{
	// 	Code: http.StatusOK,
	// 	Msg:  "success",
	// 	Data: "register success",)
}

func Login(ctx *gin.Context) {
	// 登录
}
