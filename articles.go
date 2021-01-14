package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	s "strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//Article holds text and title for the artilces with a species map.
type Article struct {
	ID    int
	Title string
	Text  string
}

//Articles checks that all species home pages articles have a map an that the map has corresponding records on the inventory API
func Articles() {
	env := GetConfig()
	db, err := sql.Open("mysql", env["USER"]+":"+env["PASSWORD"]+"@"+"/"+env["DATABASE"])
	HandleError(err)
	defer db.Close()
	query, err := db.Prepare("select ac.id, ac.title, ac.introtext from acgj_categories cat, acgj_content as ac where cat.parent_id = 279 and cat.id = ac.catid and ac.state = 1")

	HandleError(err)
	rows, err := query.Query()
	HandleError(err)

	now := time.Now().Unix()
	snow := fmt.Sprintf("%v", now)
	outputFile := "out/articulos-" + snow + ".csv"
	outPutCsv, err := os.Create(outputFile)
	HandleError(err)
	csvWriter := csv.NewWriter(outPutCsv)
	defer outPutCsv.Close()
	errorFlag := true
	for rows.Next() {
		var article Article
		err = rows.Scan(&article.ID, &article.Title, &article.Text)
		if err != nil {
			continue
		}
		errorFlag = false //could all be errors?
		re := regexp.MustCompile(`{(\s)*mapasp(\s|\w|,)*}`)
		matched := re.Find([]byte(article.Text))
		//Do we have a map plugin in the article?
		if matched != nil {
			re = regexp.MustCompile(`{(\s)*mapasp(\s)*`)
			mapSpecies := re.ReplaceAll(matched, []byte(""))
			re = regexp.MustCompile(`(\s)*}`)
			mapSpecies = re.ReplaceAll(mapSpecies, []byte(""))
			species := string(mapSpecies)
			//Do we have info for that species on the API?
			hasInfo := GetSpPoints(species)
			if !hasInfo {
				missing := make([]string, 2)
				missing[0] = species
				missing[1] = string("Mapa Sin Info en API")
				_ = csvWriter.Write(missing)
				fmt.Println(species + " Mapa sin Info en API")
			}

		} else {
			re = regexp.MustCompile(`\[i\](?P<sp1>(\w|\s)*)\[\/i\](?P<sp2>(\w|\s)*)`)
			match := FindNamedMatches(re, article.Title)
			species := fmt.Sprintf("%s %s", s.TrimSpace(match["sp1"]), s.TrimSpace(match["sp2"]))
			species = s.TrimSpace(species)
			//Can we get a species name from the title of the article?
			if species != "" {
				hasInfo := GetSpPoints(species)
				if !hasInfo {
					missing := make([]string, 2)
					missing[0] = species
					missing[1] = string("Sin Mapa y sin info en API")
					_ = csvWriter.Write(missing)
					fmt.Println(species + " Sin Mapa y sin info en API")
				} else {
					//yes,  append map to species page
					article.Text += "{mapasp " + species + "}"
					_, err := db.Query("Update acgj_content set introtext = ? where id = ?", article.Text, article.ID)
					err = nil
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println("FIXED: " + species)
					}

				}

			} else {
				missing := make([]string, 2)
				missing[0] = article.Title
				missing[1] = "Título no calza en especie"
				_ = csvWriter.Write(missing)
				fmt.Println(missing[0] + " Título no calza en especie")
			}

		}
	}
	csvWriter.Flush()
	if !errorFlag {
		fmt.Println("Output at " + outputFile)
	} else {
		fmt.Println("No valid articles to verify")
	}
}
