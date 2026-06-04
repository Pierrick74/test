package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func WriteCSV(filename string, records [][]string){
	file, err := os.Create(filename)
	if err != nil {
   		log.Fatalln("Error creating file:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	for _, record := range records {
    	if err := writer.Write(record); err != nil {
        	log.Fatalln("Error writing record to file:", err)
    	}
	}
}