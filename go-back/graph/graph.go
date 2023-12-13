package graph

import (
	"cyclesaster/models"
	"fmt"
)

func ProcessDataForGraph(data []models.DataFilters, filter2 string) interface{} {
	dataDistribution := make(map[string]int)

	fmt.Println("filter2 is ", filter2)

	for _, accident := range data {
		switch filter2 {
		case "Birth_year":
			dataDistribution[accident.Birth_year]++
		case "Department":
			if len(accident.Department) == 3 && accident.Department[2] == '0' {
				accident.Department = accident.Department[:2]
			}
			dataDistribution[accident.Department]++
		case "Year":
			dataDistribution[accident.Year]++
			// More cases to be added later
		}
	}
	return dataDistribution
}

func ProcessDataForMap(data []models.DataFilters) interface{} {
	fmt.Println("Processing data for map")
	mapData := make([]models.MapData, len(data))

	for i, accident := range data {
		fmt.Println("accident is ", accident)
		mapData[i] = models.MapData{
			Id:        accident.Id,
			Latitude:  accident.Latitude,
			Longitude: accident.Longitude,
		}
	}

	return mapData

}
