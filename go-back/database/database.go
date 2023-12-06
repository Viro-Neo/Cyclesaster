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

	query := fmt.Sprintf("SELECT a.num_acc, a.dep, u.an_nais FROM characteristics a "+
		"JOIN users u ON a.num_acc = u.num_acc "+
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