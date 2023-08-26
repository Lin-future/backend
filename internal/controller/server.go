package controller

import (
	"go-svc-tpl/api/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// >>>>>>>>>>>>>>>>>> Interface  >>>>>>>>>>>>>>>>>>
type IServerController interface {
	Link(*gin.Context, *dto.ServerLinkReq) error
	Veri(*gin.Context, *dto.ServerVeriReq) error
}

// >>>>>>>>>>>>>>>>>> Controller >>>>>>>>>>>>>>>>>
var _ IServerController = (*ServerController)(nil)

var NewServerController = func() *ServerController {
	return &ServerController{}
}

type ServerController struct {
	db *gorm.DB
}

func (c *ServerController) Link(ctx *gin.Context, req *dto.ServerLinkReq) error {
	//不是很懂这个函数要干什么
	return nil
}

func (c *ServerController) Veri(ctx *gin.Context, req *dto.ServerVeriReq) error {
	return nil
}
