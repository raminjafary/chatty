package router

import (
	"context"
	"server/adapters/handler"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitRouter(userHandler handler.UserHandler, roomHandler handler.RoomHandler, hubHandler handler.HubHandler) {
	router = gin.Default()

	router.Use(TimeoutMiddleware(5 * time.Second))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// return origin == "http://localhost:3000"
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	v1 := router.Group("/v1")

	v1.POST("/users/create", userHandler.CreateUser)
	v1.GET("/users/:id", userHandler.GetUserById)

	v1.GET("/rooms", roomHandler.GetRooms)
	v1.POST("/rooms/create", roomHandler.CreateRoom)

	v1.GET("/ws/joinRoom/:roomId", hubHandler.JoinRoom)

}

func Start(addr string) error {
	return router.Run(addr)
}

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
