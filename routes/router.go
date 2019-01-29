package routes

import (
	"online_backend/routes/api"
	"github.com/gin-gonic/gin"	
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.GET("/ping", api.Pong)

	return r
}
