package main

import (
	"database/sql"
	"fmt"
	"regexp"
	s "strings"

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

	for rows.Next() {
		var article Article
		err = rows.Scan(&article.Title, &article.Text)
		HandleError(err)
		re := regexp.MustCompile(`{(\s)*mapasp(\s|\w)*}`)
		matched := re.Find([]byte(article.Text))

		if matched != nil {
			re = regexp.MustCompile(`{(\s)*mapasp`)
			mapSpecies := re.ReplaceAll(matched, []byte(""))
			re = regexp.MustCompile(`}`)
			mapSpecies = re.ReplaceAll(mapSpecies, []byte(""))

			hasInfo := GetSpPoints("species")
			if hasInfo {
				//fmt.Println(string(mapSpecies) + "has info on the API - OK ")
			} else {
				fmt.Println(string(mapSpecies) + "doesn't has info\n")
			}

		} else {
			re = regexp.MustCompile(`\[i\](?P<sp1>(\w|\s)*)\[\/i\](?P<sp2>(\w|\s)*)`)
			match := FindNamedMatches(re, article.Title)
			fmt.Printf("%s %s has no map\n", s.TrimSpace(match["sp1"]), s.TrimSpace(match["sp2"]))

		}
	}
}
