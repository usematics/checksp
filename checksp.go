package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	s "strings"
)

type vocuher struct {
	ID               string `json:"_id"`
	HerbivoreSpecies string `json:"herbivoreSpecies"`
	CollectionDate   string `json:"collectionDate"`
	HerbivoreFamily  string `json:"herbivoreFamily"`
	Latitude         string `json:"latitude"`
	Locality         string `json:"locality"`
	Longitude        string `json:"longitude"`
	Voucher          string `json:"voucher"`
}

type spPoints struct {
	ID    string               `json:"_id"`
	Count int                  `json:"count"`
	Data  map[string][]vocuher `json:"data"`
}

type jsonResponse struct {
	Response []spPoints `json:"array"`
}

//getSpPoints checks if there is data on a given species
func getSpPoints(species string) bool {

	urlAdress := "http://localhost:3000/api/species/" + s.Replace(species, " ", "%20", -1)

	resp, err := http.Get(urlAdress)

	if resp.StatusCode != 200 {
		fmt.Println(species, "404 Not found")
		return false
	}
	HandleError(err)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading", species, "from API")
		return false
	}
	var respB []spPoints
	err = json.Unmarshal([]byte(body), &respB)
	if err != nil {
		fmt.Println("Error unmarshalling", species)
		return false
	}
	defer resp.Body.Close()

	if len(respB) > 0 {
		if respB[0].Count > 0 {
			//These will be the steps I need to go into the JSON datapoints
			// for key, oneLocation := range respB[0].Data {
			// 	fmt.Println(key, ":")
			// 	for _, dato := range oneLocation {
			// 		fmt.Println(dato.Voucher)
			// 	}
			// }
			return true
		}
	}
	return false
}

//Test script to fool around with Go.
//Concepts to explore: handling JSON responeses, csv files and http calls
//Given a csv files that has: Author, Family, Genus and Species. Checks on the local API if we have data
//Outpust a csv with species with no records
func main() {
	//	TestJSON()
	csvFile, err := os.Open("paginasptx.csv")
	HandleError(err)
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	HandleError(err)
	if _, err := os.Stat("faltan.csv"); err == nil {
		os.Remove("faltan.csv")
	}
	outPutCsv, err := os.Create("faltan.csv")
	csvWriter := csv.NewWriter(outPutCsv)
	for i, line := range csvLines {
		if i > 0 {
			//do we have a Acrotomia Mucia that should be Acrotomia mucia
			matched, _ := regexp.Match(`[A-Z](\p{L}*)\b`, []byte(line[3]))
			if matched {
				line[3] = s.ToLower(line[3])
			}
			species := s.TrimSpace(line[2]) + " " + s.TrimSpace(line[3])
			if len(species) > 0 {
				hasInfo := getSpPoints(species)
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
}

//HandleError prints and exits the program on error
func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
