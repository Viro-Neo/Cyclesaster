package database

import (
	"cyclesaster/models"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "cyclesaster_data"
)

func InitDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func FetchDataFromDB(db *sql.DB, filter1Name, filter1Value string) ([]models.DataFilters, error) {

	fmt.Println("Fetching data from DB")
	var accidents []models.DataFilters

	query := fmt.Sprintf("SELECT c.accident_id, c.department, u.birth_year FROM characteristics c "+
		"JOIN users u ON c.accident_id = u.accident_id "+
		"WHERE a.%s = $1", filter1Name)

	rows, err := db.Query(query, filter1Value)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var accident models.DataFilters
		err := rows.Scan(&accident.Id, &accident.Departement, &accident.Birth_year)
		if err != nil {
			return nil, err
		}
		accidents = append(accidents, accident)
	}

	return accidents, nil

}

func FetchFilterValuesFromDB(db *sql.DB, filterName string) ([]string, error) {
	tables := []string{"characteristics", "area", "users", "vehicles"}

	var filterValues []string

	for _, table := range tables {
		query := fmt.Sprintf("SELECT DISTINCT %s FROM %s", filterName, table)

		rows, err := db.Query(query)

		if err != nil {
			continue
		}

		for rows.Next() {
			var value sql.NullString // Change this line
			err := rows.Scan(&value)
			if err != nil {
				return nil, err
			}
			if value.Valid && !containsString(filterValues, value.String) { // Check if the value is valid before using it
				filterValues = append(filterValues, value.String)
			}
		}
		rows.Close()
	}

	return filterValues, nil
}

func containsString(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}
