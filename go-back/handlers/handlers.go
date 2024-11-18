package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"cyclesaster/database"
	"cyclesaster/graph"
	"cyclesaster/models"

	"github.com/gin-gonic/gin"
)

func GetMapData(c *gin.Context, db *sql.DB) {
	var filters []models.Filters

	for i := 1; ; i++ {
		filterName := c.Query(fmt.Sprintf("filter%d_name", i))
		filterValue := c.Query(fmt.Sprintf("filter%d_value", i))

		if filterName == "" || filterValue == "" {
			break
		}

		filters = append(filters, models.Filters{Filter_name: filterName, Filter_value: filterValue})
	}

	if len(filters) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filter1_name and filter1_value are required"})
		return
	}

	data, err := database.FetchDataFromDB(db, filters, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mapData := graph.ProcessDataForMap(data)

	c.JSON(http.StatusOK, gin.H{"data": mapData})

}

func GetGraphData(c *gin.Context, db *sql.DB) {

	var filters []models.Filters

	for i := 1; ; i++ {
		filterName := c.Query(fmt.Sprintf("filter%d_name", i))
		filterValue := c.Query(fmt.Sprintf("filter%d_value", i))

		if filterName == "" || filterValue == "" {
			break
		}

		filters = append(filters, models.Filters{Filter_name: filterName, Filter_value: filterValue})
	}

	if len(filters) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filter1_name and filter1_value are required"})
		return
	}

	perFilter := c.Query("perFilter")

	fmt.Println("perFilter is ", perFilter)

	if perFilter == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filter2 is required"})
		return
	}

	data, err := database.FetchDataFromDB(db, filters, perFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	graphData := graph.ProcessDataForGraph(data, perFilter)

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
