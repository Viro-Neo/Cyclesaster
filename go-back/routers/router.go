package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(router *gin.Engine)  {
	
	router.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "index",
		})
	})
}
