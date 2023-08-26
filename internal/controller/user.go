package controller

import (
	"go-svc-tpl/api/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// >>>>>>>>>>>>>>>>>> Interface  >>>>>>>>>>>>>>>>>>
type IUserController interface { //定义tag user 下的操作
	Register(*gin.Context, *dto.UserRegisterReq) error
	GetVeri(*gin.Context) (*dto.GetVeriResp, error)
	Login(*gin.Context, *dto.UserLoginReq) error
	GetInfo(*gin.Context) (*dto.GetUserInfoResp, error)
	ModifyInfo(*gin.Context, *dto.UserModifyInfoReq) error
	ModifyPwd(*gin.Context, *dto.UserModifyPwdReq) error
}

var _ IUserController = (*UserController)(nil)

var NewUserController = func() *UserController {
	return &UserController{}
}

type UserController struct {
	db *gorm.DB
}

// Register
func (c *UserController) Register(ctx *gin.Context, req *dto.UserRegisterReq) error {

	return nil
}

// GetVeri
func (c *UserController) GetVeri(ctx *gin.Context) (*dto.GetVeriResp, error) {
	return nil, nil
}

// Login
func (c *UserController) Login(ctx *gin.Context, req *dto.UserLoginReq) error {

	return nil

}

// GetInfo
func (c *UserController) GetInfo(ctx *gin.Context) (*dto.GetUserInfoResp, error) {

	return nil, nil

}

// ModifyInfo
func (c *UserController) ModifyInfo(ctx *gin.Context, req *dto.UserModifyInfoReq) error {

	return nil
}

// ModifyPwd
func (c *UserController) ModifyPwd(ctx *gin.Context, req *dto.UserModifyPwdReq) error {

	return nil
}
