package service

import (
	"crypto/rand"

	"landlord/config"
	"landlord/hub"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var Store sessions.Store

func init() {
	initSession()
	initRooms()
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

var emptyRooms, roomWithPassword, roomWithoutPassword, roomWithPasswordPlaying, roomWithoutPasswordPlaying *hub.Rooms

// 初始化房间
func initRooms() {
	emptyRooms.New()
	roomWithPassword.New()
	roomWithoutPassword.New()
	roomWithPasswordPlaying.New()
	roomWithoutPasswordPlaying.New()

	playConfig := config.GetPlayConfig()

	// 初始化空房间
	for i := playConfig.Rooms.Number; i > 1; i-- {
		var room = hub.Room{
			ID: i,
		}
		emptyRooms.PushBack(&room)
	}

}
