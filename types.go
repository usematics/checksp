package main

//Vocuher holds an obsevation voucher basic info
type Vocuher struct {
	ID               string `json:"_id"`
	HerbivoreSpecies string `json:"herbivoreSpecies"`
	CollectionDate   string `json:"collectionDate"`
	HerbivoreFamily  string `json:"herbivoreFamily"`
	Latitude         string `json:"latitude"`
	Locality         string `json:"locality"`
	Longitude        string `json:"longitude"`
	Voucher          string `json:"voucher"`
}

//SpPoints holds map of obseration vouchers having the location as key
type SpPoints struct {
	ID    string               `json:"_id"`
	Count int                  `json:"count"`
	Data  map[string][]Vocuher `json:"data"`
}

//JSONResponse holds the response with an array of SpPoints
type JSONResponse struct {
	Response []SpPoints `json:"array"`
}
