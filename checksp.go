package main

import (
	"fmt"
	"os"
)

//Test script to fool around with Go.
//Concepts to explore: handling JSON responeses, csv files and http calls

func main() {
	option := os.Args[1]
	switch option {
	case "json":
		TestJSON()
		break
	case "list":
		CheckList()
		break
	case "articles":
		CheckMap()
		break
	default:
		fmt.Println("No options declared")
		os.Exit(1)
	}
	os.Exit(0)
}
