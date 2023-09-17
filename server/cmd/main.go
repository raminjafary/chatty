package main

import (
	"fmt"
	"server/adapters/handler"
	"server/config"
	"server/core/hub"
	"server/core/member"
	"server/core/message"
	"server/core/room"
	"server/core/user"
	"server/logger"
	"server/router"
)

func main() {
	// configs
	config.LoadEnv()
	repo, db := config.SetupDBRepository()
	wsRepo := config.SetupWSRepository()

	db.AutoMigrate(&user.User{}, &room.Room{}, &message.Message{}, &member.Member{})

	// services
	userService := user.NewUserService(repo)
	roomService := room.NewRoomService(repo)
	hubService := hub.NewHubService(wsRepo)

	//hub
	hub := hubService.CreateHub()
	go hub.Run()

	// handlers
	userHandler := handler.NewUserHandler(userService)
	roomHandler := handler.NewRoomHandler(roomService)
	hubHandler := handler.NewHubHandler(hubService)

	logger.SetupLogger()

	// routers
	router.InitRouter(userHandler, roomHandler, hubHandler)
	router.Start(fmt.Sprintf("%s:%s", config.HttpHost, config.HttpPort))
}
