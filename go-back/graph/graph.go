package graph

import (
	"cyclesaster/models"
)

func ProcessDataForGraph(data []models.DataFilters, filter2 string) interface{} {
	dataDistribution := make(map[string]int)
	for _, accident := range data {
		switch filter2 {
		case "Birth_year":
			dataDistribution[accident.Birth_year]++
		case "Location":
			dataDistribution[accident.Departement]++
			// More cases to be added later
		}
	}
	return dataDistribution
}
