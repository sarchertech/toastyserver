package database

import (
	"database/sql"
	//blank identifer because we only care about side effects
	//from initialization not calling anything in pkg directly
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const dbName string = "ToastyTest"
const dbPath string = "./" + dbName + ".db"

//global variable for database pool
var db *sql.DB

func OpenDB() {
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

func upSchema() {
	for k, v := range schema() {
		sql := fmt.Sprintf("create table %s %s", k, v)

		_, err := db.Exec(sql)

		if err != nil {
			log.Printf("%q: %s\n", err, sql)
			return
		}
	}
}
