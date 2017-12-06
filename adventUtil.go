package advent

import (
    "path/filepath"
    "os"
    "fmt"
    "encoding/csv"
    "strconv"
)

// Read and return all of the records in a file
func ReadTsvRecordsWithDelimiter(filename string, delimiter rune) [][]string {
    absPath, _ := filepath.Abs(filename)
    file, err := os.Open(absPath)
    if err != nil {
        fmt.Println("Error reading intput file:", err)
        return nil
    }
    defer file.Close()

    reader := csv.NewReader(file)
    //Configure reader options Ref http://golang.org/src/pkg/encoding/csv/reader.go?s=#L81
    reader.Comma = delimiter           // Use tab-delimited instead of comma
    reader.Comment = '#'          //Comment character
    reader.FieldsPerRecord = -1   //Number of records per record. Set to Negative value for variable
    reader.TrimLeadingSpace = true

    records, err := reader.ReadAll()
    return records
}

func ReadTsvRecords(filename string) [][]string {
    return ReadTsvRecordsWithDelimiter(filename, '\t')
}

// Convert [][]string into [][]int to eliminate string to int conversions elsewhere
func ReadTsvRecordsToInt(filename string) [][]int {
    records := make([][]int, 0)
    strRecords := ReadTsvRecords(filename)
    for i:= 0; i < len(strRecords); i++ {
        records = append(records, make([]int, len(strRecords[i])))
        for j:=0; j < len(strRecords[i]); j++ {
            records[i][j], _ = strconv.Atoi(strRecords[i][j])  // this error handling is certainly bad
        }
    }
    return records
}
