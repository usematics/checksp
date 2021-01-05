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
	Title string
	Text  string
}

//Articles checks that all species home pages articles have a map an that the map has corresponding records on the inventory API
func Articles() {
	env := GetConfig()
	db, err := sql.Open("mysql", env["USER"]+":"+env["PASSWORD"]+"@"+"/"+env["DATABASE"])
	HandleError(err)
	defer db.Close()
	query, err := db.Prepare("select ac.title, ac.introtext from acgj_categories cat, acgj_content as ac where cat.parent_id = 279 and cat.id = ac.catid")
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

	for rows.Next() {
		var article Article
		err = rows.Scan(&article.Title, &article.Text)
		if err != nil {
			continue
		}
		re := regexp.MustCompile(`{(\s)*mapasp(\s|\w)*}`)
		matched := re.Find([]byte(article.Text))
		//Do we have a map plugin in the article?
		if matched != nil {
			re = regexp.MustCompile(`{(\s)*mapasp(\s)*`)
			mapSpecies := re.ReplaceAll(matched, []byte(""))
			re = regexp.MustCompile(`(\s)*}`)
			mapSpecies = re.ReplaceAll(mapSpecies, []byte(""))

			hasInfo := GetSpPoints("species")
			if !hasInfo {
				fmt.Println(string(mapSpecies) + " Sin Info en API")
				missing := make([]string, 2)
				missing[0] = string(mapSpecies)
				missing[1] = string("Sin Info en API")
				_ = csvWriter.Write(missing)
			}

		} else {
			re = regexp.MustCompile(`\[i\](?P<sp1>(\w|\s)*)\[\/i\](?P<sp2>(\w|\s)*)`)
			match := FindNamedMatches(re, article.Title)
			species := fmt.Sprintf("%s %s", s.TrimSpace(match["sp1"]), s.TrimSpace(match["sp2"]))
			fmt.Println(species + "Sin Mapa")
			missing := make([]string, 2)
			missing[0] = s.TrimSpace(species)
			if missing[0] == "" {
				missing[0] = article.Title
			}
			missing[1] = "Sin Mapa"
			_ = csvWriter.Write(missing)

		}
	}

	csvWriter.Flush()
}
