package main

import (
	"fmt"
	"os"

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
