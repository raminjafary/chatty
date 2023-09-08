package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func setupResponse(c *gin.Context, contentType string, body []byte, statusCode int) {
	c.Writer.Header().Set("Content-Type", contentType)

	_, err := c.Writer.Write(body)

	if err != nil {
		log.Println(err)
	}
}
