package handlers

import (
	"database/sql"
	"net/http"
	"reflect"
	"strconv"

	"cyclesaster/database"
	"cyclesaster/graph"
	"cyclesaster/models"

	"github.com/gin-gonic/gin"
)

func GetMapData(c *gin.Context, db *sql.DB) {
	filterName := c.Query("filter_name")
	filterValue := c.Query("filter_value")

	if filterName == "" || filterValue == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filter_name and filter_value are required"})
		return
	}

	data, err := database.FetchDataFromDB(db, filterName, filterValue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mapData := graph.ProcessDataForMap(data)

	c.JSON(http.StatusOK, gin.H{"data": mapData})

}

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

func GetFilters(c *gin.Context, db *sql.DB) {
	filters := models.DataFilters{}

	t := reflect.TypeOf(filters)

	var fieldNames []string

	for i := 0; i < t.NumField(); i++ {
		fieldNames = append(fieldNames, t.Field(i).Name)
	}

	c.JSON(http.StatusOK, gin.H{"filters": fieldNames})
}

func GetFiltersValues(c *gin.Context, db *sql.DB) {
	filterName := c.Query("filter_name")

	if filterName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filter_name is required"})
		return
	}

	filterValues, err := database.FetchFilterValuesFromDB(db, filterName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"filter_values": filterValues})
}

func GetAccident(c *gin.Context, db *sql.DB) {
	accidentId := c.Query("accident_id")

	if accidentId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "accident_id is required"})
		return
	}

	id, err := strconv.Atoi(accidentId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "accident_id must be an integer"})
		return
	}

	accident, err := database.FetchAccidentFromDB(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accident": accident})
}
