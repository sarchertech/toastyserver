package database

import (
	"database/sql"

	//blank identifer because we only care about side effects
	//from initialization not calling anything in pkg directly
	//"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// func FindEmployee(keyNum string) (name string, errs string) {
// 	name = ""
// 	errs = ""

// 	stmt, err := db.Prepare(`SELECT name
// 							FROM Employee
// 							Join Keyfob
// 							ON Employee.id=Keyfob.customer_id
// 							WHERE Keyfob.fob_num=?`)
// 	if err != nil {
// 		log.Println(err)
// 		errs = err.Error()
// 		return
// 	}
// 	defer stmt.Close()

// 	//cast form value to int
// 	var keyNumInt int
// 	keyNumInt, err = strconv.Atoi(keyNum) //Atoi shortcut for ParseInt(s,10,0)
// 	if err != nil {
// 		log.Println(err)
// 		errs = err.Error()
// 		return
// 	}

// 	err = stmt.QueryRow(keyNumInt).Scan(&name)
// 	if err != nil {
// 		log.Println(err)
// 		errs = err.Error()
// 		return
// 	}

// 	return
// }

func FindEmployee(keyNum int) (name string, err error) {
	var stmt *sql.Stmt
	stmt, err = db.Prepare(`SELECT name
							FROM Employee
							WHERE Employee.fob_num=?`)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(keyNum).Scan(&name)
	if err == sql.ErrNoRows {
		log.Println(err)
		err = nil
	}

	return
}

type CustomerOverview struct {
	Id     int
	Name   string
	Phone  string
	Status bool
	Level  int
}

//TODO limit results to 50
//Work on error for no rows
func RecentFiftyCustomers() (customers []CustomerOverview, err error) {
	rows, err := db.Query(`SELECT id, name, phone, status, level
						   FROM Customer`)
	if err != nil {
		return
	}
	defer rows.Close()

	//var customers []CustomerOverview

	//equivalent to while rows.Next() == true
	for rows.Next() {
		var c CustomerOverview
		rows.Scan(&c.Id, &c.Name, &c.Phone, &c.Status, &c.Level)
		//fmt.Println(customer.Name)

		customers = append(customers, c)
	}
	rows.Close()

	return
}

//TODO limit results to 50
func FindCustomersByName(name string) (customers []CustomerOverview, err error) {
	stmt, err := db.Prepare(`SELECT id, name, phone, status, level
						   	 FROM Customer
						   	 WHERE Customer.name LIKE ?`)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query("%" + name + "%")
	if err != nil {
		return
	}
	defer rows.Close()

	//equivalent to while rows.Next() == true
	for rows.Next() {
		var c CustomerOverview
		err = rows.Scan(&c.Id, &c.Name, &c.Phone, &c.Status, &c.Level)
		if err != nil {
			return
		}

		customers = append(customers, c)
	}
	if rows.Err() != nil {
		err = rows.Err()
		return
	}
	rows.Close()

	return
}
