package config

import (
	"fmt"
	"server/adapters/cache"
	"server/adapters/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *repository.DB

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
