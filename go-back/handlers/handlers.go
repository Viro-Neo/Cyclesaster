package handlers

import (
	"database/sql"
	"net/http"

	"cyclesaster/database"
	"cyclesaster/graph"

	"github.com/gin-gonic/gin"
)

func GetGraphData(c *gin.Context, db *sql.DB) {

	filter1Name := c.Query("filter1_name")
	filter1Value := c.Query("filter1_value")
	filter2 := c.Query("filter2")

	if filter2 == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filter2 is required"})
		return
	}

	data, err := database.FetchDataFromDB(db, filter1Name, filter1Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	graphData := graph.ProcessDataForGraph(data, filter2)

	c.JSON(http.StatusOK, gin.H{"data": graphData})
}
