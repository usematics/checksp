package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	s "strings"
)

//GetSpPoints checks if there is data on a given species
func GetSpPoints(species string) bool {

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
	var respB []SpPoints
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
