package controller

import (
	"go-svc-tpl/api/dto"

	"github.com/gin-gonic/gin"
)

// >>>>>>>>>>>>>>>>>> Interface  >>>>>>>>>>>>>>>>>>
type ILinkController interface { //定义tag link 下的操作
	Create(*gin.Context, *dto.LinkCreateReq) (*dto.LinkCreateResp, error)
	Delete(*gin.Context, *dto.LinkDeleteReq) error
	GetInfo(*gin.Context, *dto.GetLinkInfoReq) (*dto.GetLinkInfoResp, error)
	Update(*gin.Context, *dto.LinkUpdateReq) error
	GetList(*gin.Context, *dto.GetLinkListReq) (*dto.GetLinkListResp, error)
}
