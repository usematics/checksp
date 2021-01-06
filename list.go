package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"regexp"
	s "strings"
	"time"
)

//CheckList checks Luz María's species home page lists agains the inventory API to find species that will not have a corresponding map.
//We use this function when updating the database to spot species that could have change names and need to be updated
//Given a csv files that has: Author, Family, Genus and Species. Checks on the local API if we have data
//Outpust a csv with species with no records
func CheckList() {

	if len(os.Args) < 3 {
		HandleError(errors.New("M	issing file parameter"))
	}

	inputFile := os.Args[2]
	csvFile, err := os.Open(inputFile)
	HandleError(err)

	now := time.Now().Unix()
	snow := fmt.Sprintf("%v", now)
	outputFile := "out/articulos-" + snow + ".csv"
	fmt.Println("Checking csv file")
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	HandleError(err)
	if _, err := os.Stat(outputFile); err == nil {
		os.Remove(outputFile)
	}
	outPutCsv, err := os.Create(outputFile)
	csvWriter := csv.NewWriter(outPutCsv)
	defer outPutCsv.Close()
	for i, line := range csvLines {
		if i > 0 {
			//do we have a Acrotomia Mucia that should be Acrotomia mucia
			matched, _ := regexp.Match(`[A-Z](\p{L}*)\b`, []byte(line[3]))
			if matched {
				line[3] = s.ToLower(line[3])
			}
			species := s.TrimSpace(line[2]) + " " + s.TrimSpace(line[3])
			//species := s.TrimSpace(line[2])
			if len(species) > 0 {

				hasInfo := GetSpPoints(species)
				if !hasInfo {
					fmt.Println(line[0]+",", species)
					missing := make([]string, 3)
					missing[0] = line[0]
					missing[1] = line[1]
					missing[2] = species
					_ = csvWriter.Write(missing)
				}
			}
		}

	}
	csvWriter.Flush()
	csvFile.Close()
	fmt.Println("Done! Output in " + outputFile)
}
