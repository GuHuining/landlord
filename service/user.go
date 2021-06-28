package service

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"landlord/model"
	"net/http"
)

// Register 注册账号
func Register(c *gin.Context) {
	var request model.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		// 请求数据格式错误
		c.JSON(http.StatusBadRequest, BadRequestError{err.Error()})
		return
	}
	response, err := request.Register()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// SendValidateCode 发送验证码
func SendValidateCode(c *gin.Context) {
	var request model.SendValidateCodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, BadRequestError{err.Error()})
		return
	}
	response, err := request.SendValidateCode()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// Login 登录
func Login(c *gin.Context) {
	var request model.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, BadRequestError{err.Error()})
		return
	}
	response, err := request.Login()
	if err != nil {
		if errors.Is(err, model.PasswordNotMatchError) {
			c.JSON(http.StatusUnauthorized, response)
		} else {
			c.JSON(http.StatusInternalServerError, response)
		}
		return
	}
	// 设置session
	session := sessions.Default(c)
	session.Set("user_id", response.UserID)
	session.Set("nickname", response.Nickname)
	if err = session.Save(); err != nil {
		response.Ok = false
		response.Err = "登录失败"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

// 检查是否登录
func LoginCheck(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("user_id")
	if id == nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}