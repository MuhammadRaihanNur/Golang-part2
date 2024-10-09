package main

import (
	"encoding/csv"
	"os"
)

func main() {
	nulis := csv.NewWriter(os.Stdout)

	_ = nulis.Write([]string{"Raihan", "nivea", "hima"})
	_ = nulis.Write([]string{"budi", "hiya", "ha"})
	_ = nulis.Write([]string{"bee", "niveatu", "hah"})

	nulis.Flush()

}
