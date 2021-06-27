package service

import (
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
