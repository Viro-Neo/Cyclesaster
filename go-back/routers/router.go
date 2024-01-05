package routers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"cyclesaster/handlers"
)

func SetupRouter(db *sql.DB, router *gin.Engine) {

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "localhost:5173"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "index",
		})
	})
	router.GET("/graph", func(c *gin.Context) {
		handlers.GetGraphData(c, db)
	})
	router.GET("/get_filters", func(c *gin.Context) {
		handlers.GetFilters(c, db)
	})
	router.GET("/get_filters_values", func(c *gin.Context) {
		handlers.GetFiltersValues(c, db)
	})
	router.GET("/map", func(c *gin.Context) {
		handlers.GetMapData(c, db)
	})
	router.GET("/get_accident", func(c *gin.Context) {
		handlers.GetAccident(c, db)
	})
}
