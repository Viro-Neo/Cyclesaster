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
		case "Month":
			dataDistribution[accident.Month]++
		case "Day":
			dataDistribution[accident.Day]++
		case "Gender":
			dataDistribution[accident.Gender]++
		case "Surface":
			dataDistribution[accident.Surface]++
		case "Infrastructure":
			dataDistribution[accident.Infrastructure]++
		case "Trafic":
			dataDistribution[accident.Trafic]++
		case "Situation":
			dataDistribution[accident.Situation]++
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
