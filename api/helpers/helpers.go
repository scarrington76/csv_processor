package helpers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"slices"
	"strconv"

	"csv_processor/models"
)

func Contains(s []string, e string) bool {
	return slices.Contains(s, e)
}

func GetEnvVar(key string) string {
	if os.Getenv(key) == "" { // TODO: Change to check for env variable first
		log.Fatalf("the environment variable '%s' doesn't exist or is not set", key)
	}
	return os.Getenv(key)
}

func WriteJSON(rw http.ResponseWriter, data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		log.Println("error marshaling json: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write(js)
	if err != nil {
		log.Println("error writing json: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func ImportCSV(filepathname string) []any {
	// Open the CSV file.
	file, err := os.Open(filepathname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new CSV reader.
	reader := csv.NewReader(file)

	// Read the header row (optional).
	_, err = reader.Read() // TODO: Will I need the header record for exercise?
	if err != nil {
		log.Println("no header row found, or error reading header:", err)
	}

	var records []any
	// Read records one by one.
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break // End of file reached.
		}
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, record)

		// Print record (record[0], record[1], etc....)
		fmt.Println("Record:", record)
	}
	return records
}

func ImportCSVtoStruct(filepathname string) ([]models.CSV, error) {
	// Open the CSV file
	file, err := os.Open(filepathname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read the header row
	header, err := reader.Read()
	if err != nil {
		return nil, err
	}

	// Create a slice to store the data
	var CSV []models.CSV

	// Iterate over the rows
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Create a new Person struct
		c := models.CSV{}

		// Populate the struct based on the header
		for i, value := range row {
			switch header[i] {
			case "name":
				c.Name = value
			case "price":
				c.Price, err = strconv.ParseFloat(value, 64)
				if err != nil {
					return nil, err
				}
			case "description":
				c.Description = value
			}
		}

		// Add the struct to the slice
		CSV = append(CSV, c)
	}

	return CSV, nil
}
