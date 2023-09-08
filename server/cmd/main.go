package main

import (
	"fmt"
	"server/adapters/handler"
	"server/config"
	"server/core/room"
	"server/core/user"
	"server/logger"
	"server/router"
)

func main() {
	// configs
	config.LoadEnv()
	repo, db := config.SetupDBRepository()

	db.AutoMigrate(&user.User{}, &room.Room{})

	// services
	userService := user.NewUserService(repo)
	roomService := room.NewRoomService(repo)

	// handlers
	userHandler := handler.NewUserHandler(userService)
	roomHandler := handler.NewRoomHandler(roomService)

	logger.SetupLogger()

	// routers
	router.InitRouter(userHandler, roomHandler)
	router.Start(fmt.Sprintf("%s:%s", config.HttpHost, config.HttpPort))
}
