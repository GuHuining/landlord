package service

import (
	"crypto/rand"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

// Set global session
var Store sessions.Store

func init() {
	// Initiate session by using crypto/rand to make sure the key is random.
	key := make([]byte, 64)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	Store = cookie.NewStore(key)
}
