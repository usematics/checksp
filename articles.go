package main

import (
	"database/sql"
	"fmt"
	"regexp"
)

//Article holds text and title for the artilces with a species map.
type Article struct {
	Title string
	Text  string
}

//CheckMap checks that all species home pages have a map an that the map has corresponding records on the inventory API
func CheckMap() {
	env := GetConfig()
	db, err := sql.Open("mysql", env["USER"]+":"+env["PASSWORD"]+"@"+"/"+env["DATABASE"])
	HandleError(err)
	defer db.Close()
	query, err := db.Prepare("select ac.title, ac.introtext from acgj_categories cat, acgj_content as ac where cat.parent_id = 279 and cat.id = ac.catid and ac.introtext like \"%mapasp%\";")
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
			species := re.ReplaceAll(matched, []byte(""))
			re = regexp.MustCompile(`}`)
			species = re.ReplaceAll(species, []byte(""))
			//res := GetSpPoints("species")
			//species = string.Replace("")
			fmt.Println(string(species))
		}
	}
}
