package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ParseCSV : takes csv file name and parse it as a slice of records
func ParseCSV(fileName string) (records [][]string, err error) {
	// Open the CSV file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err = reader.ReadAll()
	return records, err
}
