package advent

import (
	"os"
	"fmt"
	"encoding/csv"
	"strconv"
	"path/filepath"
)

func CalcChecksum(filename string)(int){
	records := readRecords(filename)
	checksum := 0;
	for  i:= 0; i < len(records); i++ {
		checksum += diffLargeSmall(records[i])
	}
	return checksum
}

func CalcChecksumTwo(filename string)(int){
	records := readRecords(filename)
	checksum := 0;
	for  i:= 0; i < len(records); i++ {
		checksum += resultOfEventlyDivisible(records[i])
	}
	return checksum
}

// Return the difference between the largest and smallest converted integers in a record
func diffLargeSmall(record []string) (int) {
	smallest := 65535  		// max uint16
	largest := 0			// assuming positive integers
	position := 0
	currentValue := 0

	for (position < len(record)) {
		currentValue, _ = strconv.Atoi(record[position])
		if (currentValue < smallest) {
			smallest = currentValue;
		}
		if (currentValue > largest) {
			largest = currentValue;
		}
		position++;
	}
	return largest - smallest;
}

func resultOfEventlyDivisible(record[]string) (int) {
	for i := 0; i < len(record); i++ {
		for j:=0; j < len(record); j++ {
			if (i != j) {
				numerator, _ := strconv.Atoi(record[i])
				denominator, _ := strconv.Atoi(record[j])

				if numerator % denominator == 0 {
					return numerator / denominator;
				}
			}
		}
	}
	return -1;
}

// Read and return all of the records in a file
func readRecords(filename string) ([][]string){
	absPath, _ := filepath.Abs(filename)
	file, err := os.Open(absPath)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	//Configure reader options Ref http://golang.org/src/pkg/encoding/csv/reader.go?s=#L81
	reader.Comma = '\t' // Use tab-delimited instead of comma
	reader.Comment = '#'        //Comment character
	reader.FieldsPerRecord = -1 //Number of records per record. Set to Negative value for variable
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	return records
}
