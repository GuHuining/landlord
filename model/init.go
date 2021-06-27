package model

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"

	"landlord/config"
)

var db *sql.DB
var rdb *redis.Client

func init() {
	databaseConfig := config.GetDatabaseConfig()
	mysqlConfig := databaseConfig.Mysql
	redisConfig := databaseConfig.Redis

	// 初始化mysql连接
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", mysqlConfig.Username, mysqlConfig.Password,
		mysqlConfig.Addr, mysqlConfig.Dbname)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("init mysql: %v", err)
	}

	// 初始化redis连接
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})
}
