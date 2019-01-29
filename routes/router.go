package routes

import (
	"github.com/gin-gonic/gin"	
	_ "online_backend/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"online_backend/routes/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", api.Pong)

	return r
}
