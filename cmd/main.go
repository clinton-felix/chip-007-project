package main

import (
	"flag"
	"os"
	"github.com/clinton-felix/chip-007-project/pkg/utils"
)

func main() {

	csvFilePath := flag.String("csv", "sample.csv", "path to the csv file")
	flag.Parse()

	file, err := os.Open(*csvFilePath) // open csv file
	CheckErr(err)
	defer file.Close()

	convertCSVtoJSON(file)
	convertJSONtoCSV("output.json", "filename.output.csv")

}