package controller

import (
	"fmt"
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/dao"
	"go-svc-tpl/internal/dao/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const USER_ID_KEY = "qwq"
const LinkTable = "Links"

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
// 用gorm创建短链接
// short 如果为空字符串，则自动生成链接
// 如果起止时间为 null 则表示不设置起止时间
func (c *LinkController) Create(ctx *gin.Context, req *dto.LinkCreateReq) (*dto.LinkCreateResp, error) {
	userID := ctx.GetUint(USER_ID_KEY)
	newLink := &model.Link{
		Short:     req.Short,
		Comment:   req.Comment,
		Origin:    req.Origin,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Active:    true,
		OwnerID:   userID,
	}
	err := dao.DB(ctx).Create(newLink).Error
	if err != nil {
		return nil, err
	}

	var link model.Link
	err = dao.DB(ctx).First(&link, "Short = ?", newLink.Short).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Short is not duplicate. You can insert.")
		}
		return nil, err
	}
	return &dto.LinkCreateResp{
		Short:     newLink.Short,
		Origin:    newLink.Origin,
		Comment:   newLink.Comment,
		StartTime: newLink.StartTime,
		EndTime:   newLink.EndTime,
		Active:    newLink.Active,
	}, nil
}

// delete
func (c *LinkController) Delete(ctx *gin.Context, req *dto.LinkDeleteReq) error {
	userID := ctx.GetUint(USER_ID_KEY)
	deleteLink := &model.Link{
		Short:   req.Short,
		OwnerID: userID,
	}
	err := dao.DB(ctx).Delete(&deleteLink).Error
	if err != nil {
		return err
	}
	var link model.Link
	err = dao.DB(ctx).First(&link, "Short = ?", deleteLink.Short).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Short is not duplicate. You can insert.")
		}
		return err
	}
	return nil
}

// getinfo
func (c *LinkController) GetInfo(ctx *gin.Context, req *dto.GetLinkInfoReq) (*dto.GetLinkInfoResp, error) {
	userID := ctx.GetUint(USER_ID_KEY)
	getinfoLink := &model.Link{
		Short:   req.Short,
		OwnerID: userID,
	}

	var link model.Link
	err := dao.DB(ctx).First(&link, "Short = ?", getinfoLink.Short).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Short is not duplicate. You can insert.")
		}
		return nil, err
	}

	return &dto.GetLinkInfoResp{
		Short:     link.Short,
		Origin:    link.Origin,
		Comment:   link.Comment,
		StartTime: link.StartTime,
		EndTime:   link.EndTime,
		Active:    link.Active,
	}, nil
}

// update
func (c *LinkController) Update(ctx *gin.Context, req *dto.LinkUpdateReq) error {
	userID := ctx.GetUint(USER_ID_KEY)
	updateLink := &model.Link{
		Short:     req.Short,
		Origin:    req.Origin,
		Comment:   req.Comment,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Active:    req.Active,
		OwnerID:   userID,
	}
	err := dao.DB(ctx).Updates(&updateLink).Error
	if err != nil {
		return err
	}

	var link model.Link
	err = dao.DB(ctx).First(&link, "Short = ?", updateLink.Short).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Short is not duplicate. You can insert.")
		}
		return err
	}
	return nil
}

// getlist
func (c *LinkController) GetList(ctx *gin.Context, req *dto.GetLinkListReq) (*dto.GetLinkListResp, error) {
	var links []dto.ShortLinkModel
	var total int64

	query := dao.DB(ctx).Model(&dto.ShortLinkModel{})

	if req.PageNumber > 0 && req.PageSize > 0 {
		offset := (req.PageNumber - 1) * req.PageSize
		query = query.Offset(int(offset)).Limit(int(req.PageSize))
	}
	if result := query.Find(&links); result.Error != nil {
		return nil, result.Error
	}
	query.Count(&total)

	resp := &dto.GetLinkListResp{
		Links: links,
		Total: total,
	}

	return resp, nil
}
