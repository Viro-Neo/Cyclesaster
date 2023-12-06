package data_handling

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/lib/pq"
)

func Load_data(db *sql.DB, rootDir string) error {
	return filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".csv") {
			Load_file_data(db, path)
		}
		return nil
	})
}

func Load_file_data(db *sql.DB, filename string) {

	fmt.Println("Loading data from", filename)

	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.TrimLeadingSpace = true
	header, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	tableName, err := map_table_name(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	if tableName != "characteristics" {
		return
	}

	if !table_exists(tableName, db) {
		create_table(tableName, header, db)
	}

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	for _, record := range records {
		fmt.Println("record is", record)

		recordInterface := make([]interface{}, len(record))
		for i, v := range record {
			recordInterface[i] = v
		}

		fmt.Println("recordInterface is", recordInterface)

		placeholders := make([]string, len(record))
		for i := range placeholders {
			placeholders[i] = "$" + fmt.Sprint(i+1)
		}
		placeholderString := strings.Join(placeholders, ",")

		query := fmt.Sprintf("INSERT INTO %s VALUES (%s)", tableName, placeholderString)
		fmt.Println("query is", query)

		_, err = db.Exec(query, recordInterface...)
		if err != nil {
			fmt.Println("Error inserting row:", err)
			return
		}
	}
	fmt.Println("Data inserted successfully")
}

func map_table_name(filename string) (string, error) {
	file := filepath.Base(filename)
	fmt.Println(file)
	lowerInput := strings.ToLower(file)

	nameMap := map[string]string{
		"caracteristiques": "characteristics",
		"lieux":            "places",
		"usagers":          "users",
		"vehicules":        "vehicles",
	}

	for name, value := range nameMap {
		if strings.Contains(lowerInput, name) {
			return value, nil
		}
	}

	return "", errors.New("No table name found for " + file)
}

func table_exists(tableName string, db *sql.DB) bool {
	query := "SELECT table_name FROM information_schema.tables WHERE table_name = $1"
	var result string
	err := db.QueryRow(query, tableName).Scan(&result)
	return err == nil
}

func create_table(tableName string, columns []string, db *sql.DB) {
	query := "CREATE TABLE " + tableName + " ("
	for _, col := range columns {
		query += col + " TEXT,"
	}
	query = strings.TrimSuffix(query, ",") + ")"
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}
	fmt.Println("Table created successfully")
}
