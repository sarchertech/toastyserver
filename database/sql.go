package database

import (
	"database/sql"
	"fmt"
	//blank identifer because we only care about side effects
	_ "github.com/mattn/go-sqlite3"
	"log"
)

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

	//equivalent to while rows.Next() == true
	for rows.Next() {
		var c CustomerOverview
		rows.Scan(&c.Id, &c.Name, &c.Phone, &c.Status, &c.Level)

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

//TODO possible race conditions, check that keyfob still available and lock keyfobs
func CreateCustomer(name string, phone string, level int, keyfob int) (err error) {
	stmt, err := db.Prepare(`INSERT INTO Customer(id, name, phone, status, level, fob_num)
		                     values(?, ?, ?, ?, ?, ?)`)

	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(nil, name, phone, true, level, keyfob) //insert null into id to auto incrment
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func AvailableCustomerKeyfobs() (base10 []int32, base16 []string, err error) {
	rows, err := db.Query(`SELECT Keyfob.fob_num
						   FROM Keyfob
						   LEFT OUTER JOIN Customer
						   ON Keyfob.fob_num = Customer.fob_num
						   WHERE Customer.id IS null
						   AND Keyfob.admin = 0`)
	if err != nil {
		return
	}
	defer rows.Close()

	//equivalent to while rows.Next() == true
	for rows.Next() {
		var i int32
		err = rows.Scan(&i)
		if err != nil {
			return
		}

		base10 = append(base10, i)
		base16 = append(base16, fmt.Sprintf("%X", i))
	}
	if rows.Err() != nil {
		err = rows.Err()
		return
	}
	rows.Close()

	return
}
