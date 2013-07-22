package database

import (
	//"database/sql"

	//blank identifer because we only care about side effects
	//from initialization not calling anything in pkg directly
	//	"fmt"
	_ "github.com/mattn/go-sqlite3"

	"log"
	"strconv"
)

func FindEmployee(keyNum string) string {
	//return "Jane Butts"
	stmt, err := db.Prepare(`SELECT name
							FROM Employee
							Join Keyfob
							ON Employee.id=Keyfob.customer_id
							WHERE Keyfob.fob_num=?`)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer stmt.Close()

	//cast form value to int
	var keyNumInt int
	keyNumInt, err = strconv.Atoi(keyNum) //Atoi shortcut for ParseInt(s,10,0)
	if err != nil {
		log.Println(err)
		return ""
	}

	var name string
	err = stmt.QueryRow(keyNumInt).Scan(&name)
	if err != nil {
		log.Println(err)
		return ""
	}

	return name
}
