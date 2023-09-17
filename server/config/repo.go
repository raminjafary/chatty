package config

import (
	"fmt"
	"net/http"
	"server/adapters/cache"
	"server/adapters/repository"

	"github.com/gorilla/websocket"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *repository.DB
var upgrader *websocket.Upgrader

func SetupDBRepository() (*repository.DB, *gorm.DB) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DbHost, DbPort, DbUsername, DbPassword, DbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	redisCache, err := cache.NewRedisClient(RedisURL)

	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&user.User{}, &room.Room{})

	DB = repository.NewDB(db, redisCache)

	return DB, db
}

func SetupWSRepository() *repository.WS {
	upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return repository.NewWS(upgrader)
}
