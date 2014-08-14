package main

import (
	"flag"
	"fmt"
	"github.com/learc83/toastyserver/database"
)

func main() {
	envPtr := flag.String("env", "",
		"<production | development>, determines which db to migrate")
	flag.Parse()

	fmt.Println(*envPtr)

	switch *envPtr {
	case "production":
		createProductionDB()
		fmt.Println("Creating Production DB")
	case "development":
		createDevelopmentDB()
		fmt.Println("Creating Development DB")
	default:
		fmt.Println("No environment selected. Please pass env flag (migrate -env=development or -env=production")
	}
}

func createProductionDB() {
	err := database.DeleteDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	database.CreateAndOpenDB()
	defer database.CloseDB()

	database.UpSchema()
}
