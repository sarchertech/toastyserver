package database

import (
	//blank identifer because we only care about side effects
	//from initialization not calling anything in pkg directly
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
	"os"
	//"time"
)

const randSeed int64 = 56334

var r *rand.Rand

//WARNING -- DELETES DB
func OpenDBDevMode() {
	os.Remove(dbPath)
	OpenDB()
	upSchema()
	addDevData()
}

func addDevData() {
	//initialize random seed
	r = rand.New(rand.NewSource(randSeed))

	addFakeCustomers()
	addFakeEmployees()
	addFakeKeyfobs()
}

func addFakeCustomers() {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return
	}

	stmt, err := tx.Prepare("insert into Customer(id, name, phone, status, level) " +
		"values(?, ?, ?, ?, ?)")

	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	fakeNames := fakeNames(10)
	fakePhones := fakePhones(10)

	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(nil, fakeNames[i], fakePhones[i], 1, 3) //insert null into id to auto incrment

		if err != nil {
			log.Println(err)
			return
		}
	}
	tx.Commit()
}

func addFakeEmployees() {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return
	}

	stmt, err := tx.Prepare("insert into Employee(id, name, level) " +
		"values(?, ?, ?)")

	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	fakeNames := fakeNames(10)

	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(nil, fakeNames[i], 1) //insert null into id to auto incrment

		if err != nil {
			log.Println(err)
			return
		}
	}
	tx.Commit()
}

func addFakeKeyfobs() {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return
	}

	stmt, err := tx.Prepare("insert into Keyfob(fob_num, customer_id) " +
		"values(?, ?)")

	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	fakeNums := fakeNumbers(10)
	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(fakeNums[i], i+1)

		if err != nil {
			log.Println(err)
			return
		}
	}

	//add 5 more keyfobs with no customer associated
	moreFakeNums := fakeNumbers(10)
	for i := 0; i < 5; i++ {
		_, err = stmt.Exec(moreFakeNums[i], nil)

		if err != nil {
			log.Println(err)
			return
		}
	}

	tx.Commit()
}

func fakeNames(number int) []string {
	first := []string{"Bob", "Susanne", "Jennifer", "Georginamar", "Betty",
		"Grant", "Sarah", "Loranne", "Zorahflordian", "Seven"}

	last := []string{"Robertson", "Franklin-Louis", "Mc Van Derson-Eberts",
		"Pickles", "Wintergreen", "von Snoot", "Brown", "Smith", "CaddyWompus",
		"Cumberbatch", "Skeever-hole-tweed"}

	var nameList []string

	for i := 0; i < number; i++ {
		nameList = append(nameList, first[r.Intn(len(first))]+" "+last[r.Intn(len(last))])
	}

	return nameList
}

func fakePhones(number int) []string {
	area := []string{"770", "404", "680", "755", "804", "925"}

	var phoneList []string

	for i := 0; i < number; i++ {
		first := fmt.Sprintf("%d%d%d", r.Intn(10), r.Intn(10), r.Intn(10))
		last := fmt.Sprintf("%d%d%d%d", r.Intn(10), r.Intn(10), r.Intn(10), r.Intn(10))

		phoneList = append(phoneList, area[r.Intn(len(area))]+"-"+first+"-"+last)
	}

	return phoneList
}

//could make duplicates
func fakeNumbers(number int) []int32 {
	var numList []int32

	for i := 0; i < number; i++ {
		numList = append(numList, r.Int31())
	}

	return numList
}
