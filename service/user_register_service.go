package service

import (
	"fmt"
	"singo/model"
	"singo/serializer"

	"singo/middleware/jwtauth"

	"github.com/dgrijalva/jwt-go"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	// Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"userName" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"passwordConfirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "用户名已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() serializer.Response {
	j := &jwtauth.JWT{
		[]byte("test"),
	}
	claims := jwtauth.CustomClaims{
		1,
		service.Password,
		service.UserName,
		jwt.StandardClaims{
			ExpiresAt: 15000, //time.Now().Add(24 * time.Hour).Unix()
			Issuer:    "test",
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		fmt.Printf("fuck")
	}

	user := model.User{
		UserName: service.UserName,
		Status:   model.Active,
		Token:    token,
	}
	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}

	return serializer.BuildUserResponse(user)
}
