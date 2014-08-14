package database

import (
	"database/sql"
	"errors"
	"os"
	//blank identifer because we only care about side effects
	//from initialization not calling anything in pkg directly
	"fmt"
	_ "github.com/learc83/go-sqlite3"
	"log"
)

const dbName string = "Toasty"
const dbPath string = "./" + dbName + ".sqlite"

//global variable for database pool
var db *sql.DB

// func StartDB() {
// 	//TODO add logic to run db schema
// 	//Ueses GOENV environment variable to determine behavior
// 	env := os.Getenv("GOENV")

// 	if env == "production" {
// 		OpenDB()
// 	} else if env == "development" {
// 		OpenDBDevMode() //deletes and recreates DB
// 	}
// }

//opens only will not create a db if doesn't already exist
func OpenDB() {

	// equivalent to Python's `if not os.path.exists(filename)`
	// Exit if no db found, don't create one
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
    	log.Fatalf("no such file or directory: %s", dbPath)
    	return
	}
	
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Error on initializing database connection: %s", err.Error())
		return
	}
	//defer db.Close()

	//TODO figure out haow many max idle connections needed
	db.SetMaxIdleConns(10)

	err = db.Ping() // This makes sure the database is accessible
	if err != nil {
		log.Fatalf("Error on opening database connection: %s", err.Error())
		return
	}
}

func CreateAndOpenDB() {
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Error on initializing database connection: %s", err.Error())
		return
	}
	//defer db.Close()

	//TODO figure out haow many max idle connections needed
	db.SetMaxIdleConns(10)

	err = db.Ping() // This makes sure the database is accessible
	if err != nil {
		log.Fatalf("Error on opening database connection: %s", err.Error())
		return
	}	
}

func CloseDB() {
	db.Close()
}

func DeleteDB() (err error) {
	fmt.Print("!!!WARNING!!! Delete the database? YES or NO: ")

	var str string

	fmt.Scan(&str)
	if str != "YES" {
		err = errors.New("DB not deleted")
		return
	}

	fmt.Println("Deleting Dabase: " + dbName)
	os.Remove(dbPath)
	return
}

func UpSchema() {
	for k, v := range schema() {
		sql := fmt.Sprintf("create table %s %s", k, v)

		_, err := db.Exec(sql)

		if err != nil {
			log.Printf("%q: %s\n", err, sql)
			return
		}
	}
}
