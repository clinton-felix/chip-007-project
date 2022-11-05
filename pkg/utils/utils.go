package utils

import (
	"log"
	"encoding/json"
	"encoding/csv"
	"os"
	"fmt"
	"io/ioutil"
	"crypto/sha256"
	"strconv"
	"strings"
	"github.com/clinton-felix/chip-007-project/pkg/schema"
)

// error Check
func checkErr()  {
	if err != nil {
		log.Fatal(err)
	}
}


// hash256 function
func hash256(data []byte) string {
	hash := sha256.New()
	hash.Write([]byte(data))

	sha := hash.Sum(nil)
	return fmt.Sprintf("%x", sha)
}


// converting CSV to JSOn

func ConvertCSVtoJSON(file os.File)  {
	// dump all the metaData
	var records []MetaData

	// Reading the CSV file
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	checkErr(err)

	for i, line := range lines{
		if i == 0 {
			// skip the header line
			continue
		}

		gender := []attr1{{
			TraitType: "Gender",
			Value: line[5],
		}}

		attributes := line[6]
		var rec []attr1		// contains the appended attributes
		rec = append(rec, gender)
	}
}