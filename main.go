package main

import (
	"os"
)

// Initialize App
func main() {

	// username := os.Getenv("APP_DB_USERNAME")
	// fmt.Println(username)

	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8010")
}
