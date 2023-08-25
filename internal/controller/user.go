package controller

import (
	"go-svc-tpl/api/dto"

	"github.com/gin-gonic/gin"
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
