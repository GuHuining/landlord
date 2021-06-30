package service

import (
	"crypto/rand"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var Store sessions.Store

func init() {
	initSession()
}

// 初始化session
func initSession() {
	// 随机生成session密钥
	key := make([]byte, 64)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	Store = cookie.NewStore(key)
}
