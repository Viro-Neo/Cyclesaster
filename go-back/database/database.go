package database

import (
	"cyclesaster/models"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

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

func mapFilterNameToDBName(filterName string) string {
	if filterName == "Birth_year" {
		return "u.birth_year"
	} else if filterName == "Department" {
		return "c.department"
	} else if filterName == "Month" {
		return "c.month"
	} else if filterName == "Year" {
		return "c.year"
	} else if filterName == "Gender" {
		return "u.gender"
	} else if filterName == "Surface" {
		return "a.surface"
	} else if filterName == "Infrastructure" {
		return "a.infrastructure"
	} else if filterName == "Trafic" {
		return "a.traffic"
	} else if filterName == "Situation" {
		return "a.situation"
	}

	return ""
}

func buildDynamicQuery(filters []models.Filters) string {

	query := "SELECT c.accident_id, c.longitude, c.latitude"

	for _, filter := range filters {
		if filter.Filter_name == "Latitude" || filter.Filter_name == "Longitude" {
			continue
		}
		query += fmt.Sprintf(", %s", mapFilterNameToDBName(filter.Filter_name))
	}

	query += " FROM characteristics c JOIN users u ON c.accident_id = u.accident_id " +
		"JOIN area a ON c.accident_id = a.accident_id " +
		"JOIN vehicles v ON c.accident_id = v.accident_id WHERE "

	for i, filter := range filters {
		if i > 0 {
			query += " AND "
		}
		query += fmt.Sprintf("%s = %s", mapFilterNameToDBName(filter.Filter_name), filter.Filter_value)
	}

	/*f1 := mapFilterNameToDBName(filter1Name)
	query := fmt.Sprintf("SELECT c.accident_id, c.month, c.year, "+
		"u.birth_year, c.department, u.gender, a.surface, a.infrastructure, a.traffic, a.situation, "+
		"c.latitude, c.longitude "+
		"FROM characteristics c "+
		"JOIN users u ON c.accident_id = u.accident_id "+
		"JOIN area a ON c.accident_id = a.accident_id "+
		"JOIN vehicles v ON c.accident_id = v.accident_id "+
		"WHERE %s = $1", f1)*/

	return query
}

func scanRows(rows *sql.Rows) ([]models.DataFilters, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	var accidents []models.DataFilters

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		var accident models.DataFilters
		for i, value := range values {
			if value == nil {
				continue
			}

			switch columns[i] {
			case "id":
				accident.Id, _ = strconv.Atoi(string(value))
			case "month":
				accident.Month = string(value)
			case "year":
				accident.Year = string(value)
			case "birth_year":
				accident.Birth_year = string(value)
			case "department":
				if string(value) == "201" {
					accident.Department = "2A"
				} else if string(value) == "202" {
					accident.Department = "2B"
				} else {
					accident.Department = string(value)
				}
			case "gender":
				accident.Gender = string(value)
			case "surface":
				accident.Surface = string(value)
			case "infrastructure":
				accident.Infrastructure = string(value)
			case "traffic":
				accident.Trafic = string(value)
			case "situation":
				accident.Situation = string(value)
			case "latitude":
				accident.Latitude, _ = strconv.ParseFloat(string(value), 64)
			case "longitude":
				accident.Longitude, _ = strconv.ParseFloat(string(value), 64)
			}
		}

		accidents = append(accidents, accident)
	}

	return accidents, nil
}

/*func scanRows(rows *sql.Rows) ([]models.DataFilters, error) {
	var accidents []models.DataFilters

	for rows.Next() {
		var accident models.DataFilters
		var id, month, year, birth_year, department, gender, surface, infrastructure, trafic, situation sql.NullString
		var latitude, longitude sql.NullFloat64
		err := rows.Scan(&id, &month, &year, &birth_year, &department, &gender,
			&surface, &infrastructure, &trafic, &situation, &latitude, &longitude)
		if err != nil {
			fmt.Println("error is ", err)
			return nil, err
		}

		if id.Valid {
			accident.Id, _ = strconv.Atoi(id.String)
		}
		if month.Valid {
			accident.Month = month.String
		}
		if year.Valid {
			accident.Year = year.String
		}
		if birth_year.Valid {
			accident.Birth_year = birth_year.String
		}
		if department.Valid {
			if department.String == "201" {
				accident.Department = "2A"
			} else if department.String == "202" {
				accident.Department = "2B"
			} else {
				accident.Department = department.String
			}
		}
		if gender.Valid {
			accident.Gender = gender.String
		}
		if surface.Valid {
			accident.Surface = surface.String
		}
		if infrastructure.Valid {
			accident.Infrastructure = infrastructure.String
		}
		if trafic.Valid {
			accident.Trafic = trafic.String
		}
		if latitude.Valid {
			accident.Latitude = latitude.Float64
		}
		if longitude.Valid {
			accident.Longitude = longitude.Float64
		}
		if situation.Valid {
			accident.Situation = situation.String
		}

		fmt.Println("adding accident ", accident)
		accidents = append(accidents, accident)
	}

	return accidents, nil
}*/

func FetchDataFromDB(db *sql.DB, filters []models.Filters) ([]models.DataFilters, error) {

	fmt.Println("Fetching data from DB")
	var accidents []models.DataFilters

	query := buildDynamicQuery(filters)

	fmt.Println("query is ", query)

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fmt.Println("rows are ", rows)

	accidents, _ = scanRows(rows)

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
			var value sql.NullString
			err := rows.Scan(&value)
			if err != nil {
				return nil, err
			}
			if value.Valid && !containsString(filterValues, value.String) {
				if !strings.HasPrefix(value.String, "0") {
					originalstr := "0" + value.String
					if !containsString(filterValues, originalstr) {
						filterValues = append(filterValues, value.String)
					}
				} else {
					filterValues = append(filterValues, value.String)
				}
			}
		}
		rows.Close()
	}

	return filterValues, nil
}

func FetchAccidentFromDB(db *sql.DB, accident_id int) (models.DataFilters, error) {
	var accident models.DataFilters

	query := fmt.Sprintf("SELECT c.accident_id, c.month, c.year, "+
		"u.birth_year, c.department, u.gender, a.surface, a.infrastructure, a.traffic, a.situation, "+
		"c.latitude, c.longitude "+
		"FROM characteristics c "+
		"JOIN users u ON c.accident_id = u.accident_id "+
		"JOIN area a ON c.accident_id = a.accident_id "+
		"JOIN vehicles v ON c.accident_id = v.accident_id "+
		"WHERE c.accident_id = %d", accident_id)

	fmt.Println("query is ", query)

	row, err := db.Query(query)

	if err != nil {
		return accident, err
	}
	defer row.Close()

	accidents, _ := scanRows(row)

	if len(accidents) > 0 {
		accident = accidents[0]
	}

	return accident, nil
}

func containsString(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}
