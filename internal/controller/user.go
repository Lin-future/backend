package controller

import (
	"fmt"
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/dao"
	"go-svc-tpl/internal/dao/model"
	"regexp"

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
	userID := ctx.GetUint(USER_ID_KEY)
	//验证邮箱格式
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	_, err := regexp.MatchString(emailPattern, req.Email)
	if err != nil {
		fmt.Println("wrong email pattern", err)
		return err
	}
	newUser := &model.User{
		ID:       userID,
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}
	err = dao.DB(ctx).Create(newUser).Error
	if err != nil {
		return err
	}
	//验证邮箱唯一性
	var user model.User
	err = dao.DB(ctx).First(&user, "Email = ?", newUser.Email).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Email is not duplicate. You can insert.")
		}
		return err
	}
	return nil
}

// GetVeri
func (c *UserController) GetVeri(ctx *gin.Context) (*dto.GetVeriResp, error) {
	//no
	return nil, nil
}

// Login
func (c *UserController) Login(ctx *gin.Context, req *dto.UserLoginReq) error {
	//no
	return nil

}

// GetInfo
func (c *UserController) GetInfo(ctx *gin.Context) (*dto.GetUserInfoResp, error) {
	userID := ctx.GetUint(USER_ID_KEY)
	var user model.User
	err := dao.DB(ctx).First(&user, "ID = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return &dto.GetUserInfoResp{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// ModifyInfo
func (c *UserController) ModifyInfo(ctx *gin.Context, req *dto.UserModifyInfoReq) error {
	userID := ctx.GetUint(USER_ID_KEY)
	//验证邮箱格式
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	_, err := regexp.MatchString(emailPattern, req.Email)
	if err != nil {
		fmt.Println("wrong email pattern", err)
		return err
	}
	newUser := &model.User{
		ID:    userID,
		Email: req.Email,
		Name:  req.Name,
	}
	err = dao.DB(ctx).Updates(&newUser).Error
	if err != nil {
		return err
	}
	return nil
}

// ModifyPwd
func (c *UserController) ModifyPwd(ctx *gin.Context, req *dto.UserModifyPwdReq) error {
	userID := ctx.GetUint(USER_ID_KEY)
	var user model.User
	err := dao.DB(ctx).First(&user, "ID = ?", userID).Error
	if err != nil {
		return err
	}
	if user.Password != req.OldPwd {
		fmt.Println("wrong password")
	}
	newUser := &model.User{
		Password: req.NewPwd,
	}
	err = dao.DB(ctx).Updates(&newUser).Error
	if err != nil {
		return err
	}
	return nil
}
