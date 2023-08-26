package controller

import (
	"go-svc-tpl/api/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// >>>>>>>>>>>>>>>>>> Interface  >>>>>>>>>>>>>>>>>>
type ILinkController interface { //定义tag link 下的操作
	Create(*gin.Context, *dto.LinkCreateReq) (*dto.LinkCreateResp, error)
	Delete(*gin.Context, *dto.LinkDeleteReq) error
	GetInfo(*gin.Context, *dto.GetLinkInfoReq) (*dto.GetLinkInfoResp, error)
	Update(*gin.Context, *dto.LinkUpdateReq) error
	GetList(*gin.Context, *dto.GetLinkListReq) (*dto.GetLinkListResp, error)
}

// >>>>>>>>>>>>>>>>>> Controller >>>>>>>>>>>>>>>>>>

// check interface implementation
var _ ILinkController = (*LinkController)(nil)

var NewLinkController = func() *LinkController {
	return &LinkController{}
}

type LinkController struct {
	db *gorm.DB
}

// create
func (c *LinkController) Create(ctx *gin.Context, req *dto.LinkCreateReq) (*dto.LinkCreateResp, error) {
	return nil, nil
}

// delete
func (c *LinkController) Delete(ctx *gin.Context, req *dto.LinkDeleteReq) error {
	return nil
}

// getinfo
func (c *LinkController) GetInfo(ctx *gin.Context, req *dto.GetLinkInfoReq) (*dto.GetLinkInfoResp, error) {
	return nil, nil
}

// update
func (c *LinkController) Update(ctx *gin.Context, req *dto.LinkUpdateReq) error {
	return nil
}

// getlist
func (c *LinkController) GetList(ctx *gin.Context, req *dto.GetLinkListReq) (*dto.GetLinkListResp, error) {
	return nil, nil
}
