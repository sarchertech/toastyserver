package main

import (
	"github.com/learc83/toastyserver/database"
)

func main() {
	createDevelopmentDB()
}

func createProductionDB() {
	database.DeleteDB()
	database.OpenDB()
	defer database.CloseDB()

	database.UpSchema()
}
