package controller

import (
	"go-svc-tpl/api/dto"

	"github.com/gin-gonic/gin"
)

// >>>>>>>>>>>>>>>>>> Interface  >>>>>>>>>>>>>>>>>>
type IServerController interface {
	Link(*gin.Context, *dto.ServerLinkReq) error
	Veri(*gin.Context, *dto.ServerVeriReq) error
}
