package models

type DataFilters struct {
	Id             int
	Day            string
	Month          string
	Year           string
	Birth_year     string
	Department     string
	Gender         string
	Surface        string
	Infrastructure string
	Trafic         string
	Situation      string
	Latitude       float64
	Longitude      float64
}

type MapData struct {
	Id        int
	Latitude  float64
	Longitude float64
}
