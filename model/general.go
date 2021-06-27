package model

import "errors"

const (
	DuplicatedCode = 1062
)

var DuplicatedError = errors.New("账号已存在")
var ValidateCodeError = errors.New("验证码错误")
var SendMailFailedError = errors.New("发送邮件失败")
var RequestTooQuick = errors.New("请求过快")
var PasswordNotMatchError = errors.New("用户名或密码不正确")