package service

import (
	"fmt"
	"riji/model"
	. "riji/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (s *RijiServer) UserLogin(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBind(&req); err != nil {
		ValidateError(c, ParamError, "请求参数解析失败")
		return
	}
	if req.Password == "" || req.UserName == "" {
		ValidateError(c, ParamError, "缺少必要参数")
		return
	}
	var user model.User
	user.UserName = req.UserName
	if err := s.dao.GetUserByUserName(&user); err != nil {
		ValidateError(c, UserNotFound, "用户信息不存在")
		return
	}
	// 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		ValidateError(c, UserPasswordError, fmt.Sprintf("用户密码错误 %s", err.Error()))
		return
	}
	// 设置session
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Set("user_name", user.Name)
	session.Save()
	Response(c, map[string]interface{}{"name": user.Name, "id": user.ID})
}

func (s *RijiServer) UserRegister(c *gin.Context) {
	var req RegisterReq
	if err := c.ShouldBind(&req); err != nil {
		ValidateError(c, ParamError, "请求参数解析失败")
		return
	}
	if req.Password == "" || req.UserName == "" || req.Email == "" || req.Name == "" {
		ValidateError(c, ParamError, "缺少必要参数")
		return
	}
	var user = model.User{
		LastLoginTime: nil,
		Email:         req.Email,
		Name:          req.Name,
		LastLoginIp:   c.ClientIP(),
	}
	// 检查账号名是否重复
	user.UserName = req.UserName
	if !s.dao.CheckUserName(req.UserName) {
		ValidateError(c, NormalError, "用户名重复")
		return
	}
	// 加密密码
	sePass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 7)
	if err != nil {
		ValidateError(c, NormalError, "加密密码失败")
		return
	}
	user.Password = string(sePass)
	if err := s.dao.CreateTable(&user); err != nil {
		ValidateError(c, DBError, fmt.Sprintf("新增记录失败: %s", err.Error()))
		return
	}

	// 设置session
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Set("user_name", user.Name)
	session.Save()
	Response(c, map[string]interface{}{"name": user.Name, "id": user.ID})
}

func (s *RijiServer) UserLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	Response(c, "登出成功")
}
