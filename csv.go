package nutils

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func readCsv(fileName string, startIndex, numberOfColumns int) [][]interface{} {
	// Open the file
	max := false
	csvfile, err := os.Open(fileName)
	data := make([][]interface{}, 0)
	if err != nil {
		fmt.Println("Couldn't open the csv file", err)
		return data
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	if numberOfColumns == -1 {
		max = true
	}

	// Iterate through the records
	for {
		data1 := make([]interface{}, 0)
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return data
		}
		if max {
			numberOfColumns = len(record)
		}
		for i := startIndex; i < numberOfColumns; i++ {
			data1 = append(data1, record[i])
		}
		data = append(data, data1)
	}
	return data
}
