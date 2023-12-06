package routers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"cyclesaster/handlers"
)

func SetupRouter(db *sql.DB, router *gin.Engine) {

	router.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "index",
		})
	})
	router.GET("/graph", func(c *gin.Context) {
		handlers.GetGraphData(c, db)
	})
}