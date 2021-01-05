package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

//HandleError prints and exits the program on error
func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//GetConfig reurns a map[string]string with the enviroment variables
func GetConfig() map[string]string {
	var env map[string]string
	env, err := godotenv.Read()
	HandleError(err)
	return env
}

//FindNamedMatches returns a map with the named matches for a regexp
func FindNamedMatches(regex *regexp.Regexp, str string) map[string]string {
	match := regex.FindStringSubmatch(str)

	results := map[string]string{}
	for i, name := range match {
		results[regex.SubexpNames()[i]] = name
	}
	return results
}
