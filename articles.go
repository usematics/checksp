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
	ID     int
	Title  string
	Text   string
	Author string
}

//Articles checks that all species home pages articles have a map an that the map has corresponding records on the inventory API
func Articles() {
	env := GetConfig()
	db, err := sql.Open("mysql", env["USER"]+":"+env["PASSWORD"]+"@"+"/"+env["DATABASE"])
	HandleError(err)
	defer db.Close()
	query, err := db.Prepare("select ac.id, ac.title, ac.introtext, u.name from acgj_categories cat, acgj_content as ac, acgj_users as u where cat.parent_id = 279 and cat.id = ac.catid and ac.state = 1 and ac.created_by = u.id")

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
		err = rows.Scan(&article.ID, &article.Title, &article.Text, &article.Author)
		if err != nil {
			continue
		}
		errorFlag = false //could all be errors?
		re := regexp.MustCompile(`{(\s)*mapasp(\s|\w|,)*}`)
		matched := re.Find([]byte(article.Text))
		var message string
		//Do we have a map plugin in the article?
		if matched != nil {
			re = regexp.MustCompile(`{(\s)*mapasp(\s)*`)
			mapSpecies := re.ReplaceAll(matched, []byte(""))
			re = regexp.MustCompile(`(\s)*}`)
			mapSpecies = re.ReplaceAll(mapSpecies, []byte(""))
			species := string(mapSpecies)
			article.Title = species
			//Do we have info for that species on the API?
			hasInfo := GetSpPoints(species)
			if !hasInfo {
				message = "Mapa Sin Info en API"
			}

		} else {
			re = regexp.MustCompile(`\[i\](?P<sp1>(\w|\s)*)\[\/i\](?P<sp2>(\w|\s)*)`)
			match := FindNamedMatches(re, article.Title)
			species := fmt.Sprintf("%s %s", s.TrimSpace(match["sp1"]), s.TrimSpace(match["sp2"]))
			species = s.TrimSpace(species)

			//Can we get a species name from the title of the article?
			if species != "" {
				article.Title = species
				hasInfo := GetSpPoints(species)
				if !hasInfo {
					message = "Sin Mapa y sin info en API"
				} else {
					//yes,  append map to species page
					article.Text += "{mapasp " + species + "}"
					_, err := db.Query("Update acgj_content set introtext = ? where id = ?", article.Text, article.ID)
					if err != nil {
						message = err.Error()
					} else {
						message = "Arreglada"
					}

				}

			} else {
				message = "Título no calza en especie"
			}

		}
		if message != "" {
			logInCsv(article, message, csvWriter)
		}
	}
	csvWriter.Flush()
	if !errorFlag {
		fmt.Println("Output at " + outputFile)
	} else {
		fmt.Println("No valid articles to verify")
	}
}

func logInCsv(article Article, message string, csvWriter *csv.Writer) {
	missing := make([]string, 3)
	missing[0] = article.Author
	missing[1] = article.Title
	missing[2] = message
	_ = csvWriter.Write(missing)
	fmt.Println(missing[1] + " " + message)
}
