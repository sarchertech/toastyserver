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
		err := database.DeleteDB()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Creating Production DB")
		createProductionDB()
	case "development":
		err := database.DeleteDB()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Creating Development DB")
		createDevelopmentDB()
	default:
		fmt.Println("No environment selected. Please pass env flag (migrate -env=development or -env=production")
	}
}

func createProductionDB() {
	database.CreateAndOpenDB()
	defer database.CloseDB()

	database.UpSchema()
}
