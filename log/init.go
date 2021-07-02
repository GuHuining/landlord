package log

import (
	"log"
	"os"
)

var MyLog *log.Logger

func init() {
	f, err := os.OpenFile("./landlord.log", os.O_CREATE|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		log.Fatalf("init log: %v", err)
	}
	MyLog = log.New(f, "mylog", log.LstdFlags)
}