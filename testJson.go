package main

import (
	"fmt"
	"os"
)

//TestJSON helps me reverse engeneering the JSON format
//it was good to understand how to shape the struct chains
func TestJSON() {
	voucher1 := Vocuher{
		ID:               "5fcfb3393f2e48e107adc0a7",
		HerbivoreSpecies: "Zaretis crawfordhilli",
		CollectionDate:   "10/22/2004",
		HerbivoreFamily:  "Nymphalidae",
		Latitude:         "10.98680",
		Locality:         "Sendero Evangelista",
		Longitude:        "-85.42083",
		Voucher:          "04-SRNP-55848",
	}

	voucher2 := Vocuher{
		ID:               "5f3333393f2e48e107adc0a7",
		HerbivoreSpecies: "Zaretis crawfordhilli",
		CollectionDate:   "10/22/2004",
		HerbivoreFamily:  "Nymphalidae",
		Latitude:         "10.98680",
		Locality:         "Sendero Evangelista",
		Longitude:        "-85.42083",
		Voucher:          "05-SRNP-55848",
	}

	locals := []Vocuher{}
	locals = append(locals, voucher1, voucher2)
	//decare and initialize the map with data, to initialize it empty:
	//localities := make(map[string][]vocuher)
	localities := map[string][]Vocuher{
		"Sendero Evangelista": locals,
	}

	responseB := []SpPoints{} ///
	spPoint := SpPoints{
		ID:    "ZC",
		Count: 180.0,
		Data:  localities,
	}

	responseB = append(responseB, spPoint)

	fmt.Println(spPoint)
	os.Exit(0)
}
