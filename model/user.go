package model

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"landlord/tools"
	"math/rand"
	"time"
)

type RegisterRequest struct {
	Username     string `json:"username" binding:"required,max=20,min=6"`
	Password     string `json:"password" binding:"required,max=20,min=6"`
	Nickname     string `json:"nickname" binding:"required,max=20"`
	Email        string `json:"email" binding:"required,email,max=254"`
	ValidateCode string `json:"validate_code" binding:"required,len=6"`
}

type RegisterResponse struct {
	Ok  bool   `json:"ok"`
	Err string `json:"err"`
}

// Register 用户注册
func (request *RegisterRequest) Register() (response RegisterResponse, err error) {
	// 判断验证码是否相符
	result, err := rdb.Get("validate_code:" + request.Email).Result()
	if err != nil || result != request.ValidateCode {
		response.Err = ValidateCodeError.Error()
		err = ValidateCodeError
		return
	}
	rdb.Del("validate_code:" + request.Email)

	// 加密密码,并返回加密后的密码和盐值
	encodedPassword, salt := tools.Md5EncodingPassword(request.Password)
	// 插入登录信息
	_, err = db.Exec("INSERT INTO user (username, password, salt, nickname, mail) VALUE(?, ?, ?, ?, ?)",
		request.Username, encodedPassword, salt, request.Nickname, request.Email)
	if err != nil {
		if err.(*mysql.MySQLError).Number == DuplicatedCode { // 已有该账号
			err = DuplicatedError
		}
		response.Err = err.Error()
		return
	}
	response.Ok = true
	return
}

type SendValidateCodeRequest struct {
	Email string `json:"email" binding:"email,max=254"`
}

type SendValidateCodeResponse struct {
	Ok  bool   `json:"ok"`
	Err string `json:"err"`
}

// 发送验证码
func (request *SendValidateCodeRequest) SendValidateCode() (response SendValidateCodeResponse, err error) {
	// 判断生成验证码间隔是否小于1分钟
	result, err := rdb.TTL("validate_time:" + request.Email).Result()
	if result > 0 {
		err = RequestTooQuick
		response.Err = RequestTooQuick.Error()
		return
	}
	// 随机生成验证码
	validateCode := rand.Intn(900000) + 100000
	err = rdb.Set("validate_code:"+request.Email, validateCode, 15*time.Minute).Err()
	if err != nil {
		return
	}
	rdb.Set("validate_time:"+request.Email, 1, time.Minute) // 设置获取验证码间隔

	// 发送邮件
	template := `
<h1>欢迎注册</h1>
<p>验证码: %d
`
	body := fmt.Sprintf(template, validateCode)
	err = tools.SendMail(request.Email, "验证码", body)
	if err != nil {
		response.Err = SendMailFailedError.Error()
		return
	}
	response.Ok = true
	return
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=6,max=20"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

type LoginResponse struct {
	Ok       bool   `json:"ok"`
	Err      string `json:"err"`
	UserID   int    `json:"user_id"`
	Nickname string `json:"nickname"`
}

// Login 登录
func (request *LoginRequest) Login() (response LoginResponse, err error) {
	var password, salt, nickname string
	var userID int
	err = db.QueryRow("SELECT user_id, nickname, password, salt FROM user WHERE username=?", request.Username).
		Scan(&userID, &nickname, &password, &salt)
	if err != nil {
		response.Err = PasswordNotMatchError.Error()
		return
	}
	if !tools.ValidatePassword(password, request.Password, salt) { // 用户名与密码不匹配
		response.Err = PasswordNotMatchError.Error()
		err = PasswordNotMatchError
		return
	}
	response.Nickname = nickname
	response.UserID = userID
	response.Ok = true
	return
}
