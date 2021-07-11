package infrastructure

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitGinRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
