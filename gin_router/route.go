package gin_router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	v1 := router.Group("admin")
	{
	}
	router.Run(":8080")
}
