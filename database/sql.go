package database

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	//blank identifer because we only care about side effects
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//TODO abstract out with FindRecord just like CreateRecord
func FindEmployee(keyNum int) (name string, err error) {
	var stmt *sql.Stmt
	stmt, err = db.Prepare(`SELECT Name
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

//TODO limit results to 50
//Work on error for no rows
//TODO abstract out with ListRecords just like CreateRecord
func RecentFiftyCustomers() (customers []Customer, err error) {
	rows, err := db.Query(`SELECT Id, Name, Phone, Status, Level
						   FROM Customer`)
	if err != nil {
		return
	}
	defer rows.Close()

	//equivalent to while rows.Next() == true
	for rows.Next() {
		var c Customer
		rows.Scan(&c.Id, &c.Name, &c.Phone, &c.Status, &c.Level)

		customers = append(customers, c)
	}
	rows.Close()

	return
}

//TODO limit results to 50
//TODO abstract out with ListRecords just like CreateRecord
func FindCustomersByName(name string) (customers []Customer, err error) {
	stmt, err := db.Prepare(`SELECT Id, Name, Phone, Status, Level
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
		var c Customer
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

//creates a record from an initalized struct, set autoIncrement to true if the
//first field defined in the struct is an autoincrement field
//Uses reflection to set the Table name to the Type name of the struct, and to get
//the names and values of an arbitrary number of fields
//TODO check for race condition when adding new customer--make sure keyfob exists
func CreateRecord(record interface{}, autoIncrement bool) (err error) {
	t := reflect.TypeOf(record)
	v := reflect.ValueOf(record)

	var fields []string
	var values []interface{}

	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i).Name)
		values = append(values, v.Field(i).Interface())
	}

	fieldStr := strings.Join(fields, ", ")
	qMarks := strings.Repeat("?,", t.NumField()-1) + "?"

	sqls := fmt.Sprintf(`INSERT INTO %s(%s)
		                 values(%s)`, t.Name(), fieldStr, qMarks)
	log.Println(sqls)

	stmt, err := db.Prepare(sqls)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	if autoIncrement {
		values[0] = nil //set first value to nil so sqlite will autoincrement
	}

	_, err = stmt.Exec(values...)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func AvailableCustomerKeyfobs() (base10 []int32, base16 []string, err error) {
	rows, err := db.Query(`SELECT Keyfob.Fob_num
						   FROM Keyfob
						   LEFT OUTER JOIN Customer
						   ON Keyfob.fob_num = Customer.Fob_num
						   WHERE Customer.Id IS null
						   AND Keyfob.Admin = 0`)
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
