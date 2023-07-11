package database

import "os"

func Get_connection_string() string {
	filePath := "database/startup/connection_string.txt"
	readFile, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(readFile)
}
