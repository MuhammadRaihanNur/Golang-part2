package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "raihan,muhammad,nur\n" +
		"budi,muhammad,nur\n" +
		"joko,arif,plus"

	csvReader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}
