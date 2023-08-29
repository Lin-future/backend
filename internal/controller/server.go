package controller

import (
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/dao/model"
	"net/http"

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

// 重定向到相应链接
// 当出现问题时，会按标准 response 返回，由前端显示相应页面
func (c *ServerController) Link(ctx *gin.Context, req *dto.ServerLinkReq) error {
	shortLink := req.Short

	// 在数据库中查找对应的长链接
	var link model.Link
	if err := c.db.Table(model.LinkTable).Where("short = ?", shortLink).First(&link).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 如果找不到对应的长链接，返回404错误
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		} else {
			// 如果发生其他错误，返回500错误
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return err
	}

	// 如果找到了对应的长链接，进行重定向
	ctx.Redirect(http.StatusMovedPermanently, link.Origin)
	return nil
}

func (c *ServerController) Veri(ctx *gin.Context, req *dto.ServerVeriReq) error {
	//no
	return nil
}
