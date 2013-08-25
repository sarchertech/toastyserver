package database

import (
	//blank identifer because we only care about side effects
	//from initialization not calling anything in pkg directly
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
	"os"
	"time"
)

const randSeed int64 = 56355145

var r *rand.Rand

// func OpenProduction() {
// 	//os.Remove(dbPath)
// 	OpenDB()
// 	//upSchema()

// 	// bed1 := Bed{Bed_num: 1, Level: 1, Max_time: 15, Name: "SunDash 232"}
// 	// bed2 := Bed{Bed_num: 2, Level: 2, Max_time: 12, Name: "Standup"}
// 	// bed3 := Bed{Bed_num: 3, Level: 1, Max_time: 15, Name: "SunDash 232"}
// 	// bed4 := Bed{Bed_num: 4, Level: 3, Max_time: 12, Name: "Ergoline 600"}
// 	// bed5 := Bed{Bed_num: 5, Level: 1, Max_time: 15, Name: "SunDash 232"}
// 	// bed6 := Bed{Bed_num: 6, Level: 2, Max_time: 15, Name: "Heartland"}

// 	// CreateRecord(bed1)
// 	// CreateRecord(bed2)
// 	// CreateRecord(bed3)
// 	// CreateRecord(bed4)
// 	// CreateRecord(bed5)
// 	// CreateRecord(bed6)

// 	// keyfob2 := Keyfob{Fob_num: 9873, Admin: false}
// 	// CreateRecord(keyfob2)

// 	// customer := Customer{Name: "Jane Tanner", Level: 3, Fob_num: 9873,
// 	// 	Phone: "770-949-1622", Status: true}
// 	// CreateRecord(customer)

// }

//WARNING -- DELETES DB
func OpenDBDevMode() {
	os.Remove(dbPath)
	OpenDB()
	UpSchema()
	addDevData()

	keyfob := Keyfob{Fob_num: 12107728, Admin: true}
	CreateRecord(keyfob)

	employee := Employee{Name: "Seth", Level: 3, Fob_num: 12107728}
	CreateRecord(employee)

	keyfob2 := Keyfob{Fob_num: 9873, Admin: false}
	CreateRecord(keyfob2)

	customer := Customer{Name: "Jane Tanner", Level: 3, Fob_num: 9873,
		Phone: "770-949-1622", Status: true}
	CreateRecord(customer)

	customer2 := Customer{Name: "Fred Tanner", Level: 3, Fob_num: 9871,
		Phone: "770-949-1622", Status: false}
	CreateRecord(customer2)

	session := Session{
		Bed_num:     5,
		Customer_id: 11,
		Time_stamp:  time.Now().Unix() - 43201}

	err := CreateRecord(session)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("bed started, session created")

	//time.Sleep(1.5 * 1e9)

	session.Time_stamp = time.Now().Unix() - 50000

	err = CreateRecord(session)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("bed started, session created")
}

func addDevData() {
	//initialize random seed
	r = rand.New(rand.NewSource(randSeed))

	fakeKeyNums := addFakeKeyfobs()
	addFakeCustomers(fakeKeyNums)
	addFakeEmployees(fakeKeyNums)
	addFakeBeds()
}

func addFakeCustomers(fakeKeyNums []int32) {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return
	}

	stmt, err := tx.Prepare("insert into Customer(id, name, phone, status, level, fob_num) " +
		"values(?, ?, ?, ?, ?, ?)")

	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	fakeNames := fakeNames(10)
	fakePhones := fakePhones(10)

	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(nil, fakeNames[i], fakePhones[i], 1, r.Intn(5)+1, fakeKeyNums[i]) //insert null into id to auto incrment

		if err != nil {
			log.Println(err)
			return
		}
	}
	tx.Commit()
}

func addFakeEmployees(fakeKeyNums []int32) {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return
	}

	stmt, err := tx.Prepare("insert into Employee(id, name, level, fob_num) " +
		"values(?, ?, ?, ?)")

	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	fakeNames := fakeNames(10)

	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(nil, fakeNames[i], 1, fakeKeyNums[i+10]) //insert null into id to auto incrment

		if err != nil {
			log.Println(err)
			return
		}
	}
	tx.Commit()
}

func addFakeKeyfobs() (fakeNums []int32) {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return
	}

	stmt, err := tx.Prepare("insert into Keyfob(fob_num, admin) " +
		"values(?, ?)")

	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	fakeNums = fakeNumbers(20)

	//fake Customer Keyfobs
	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(fakeNums[i], false)

		if err != nil {
			log.Println(err)
			return
		}
	}
	//fake Employee Keyfobs
	for i := 10; i < 20; i++ {
		_, err = stmt.Exec(fakeNums[i], true)

		if err != nil {
			log.Println(err)
			return
		}
	}

	//add 5 more keyfobs with no customer associated
	moreFakeNums := fakeNumbers(10)
	for i := 0; i < 5; i++ {
		_, err = stmt.Exec(moreFakeNums[i], false)

		if err != nil {
			log.Println(err)
			return
		}
	}

	//add 5 more keyfobs with no  employee associated
	for i := 5; i < 10; i++ {
		_, err = stmt.Exec(moreFakeNums[i], true)

		if err != nil {
			log.Println(err)
			return
		}
	}

	tx.Commit()

	return
}

func addFakeBeds() {
	//Need to fix numbering, quick hax
	for i := 0; i < 10; i++ {
		bed := Bed{Bed_num: i + 1, Level: 1, Max_time: 15, Name: "Sundash 232"}
		CreateRecord(bed)
	}

	for i := 0; i < 5; i++ {
		bed := Bed{Bed_num: i + 11, Level: 2, Max_time: 12, Name: "Ameribed 64"}
		CreateRecord(bed)
	}

	for i := 0; i < 2; i++ {
		bed := Bed{Bed_num: i + 16, Level: 3, Max_time: 10, Name: "Bad Ass Bed"}
		CreateRecord(bed)
	}
}

func fakeNames(number int) []string {
	first := []string{"Bob", "Susanne", "Jennifer", "Georginamar", "Betty",
		"Grant", "Sarah", "Loranne", "Zorahflordian", "Seven"}

	last := []string{"Robertson", "Franklin-Louis", "Mc Van Derson-Eberts",
		"Pickles", "Wintergreen", "von Snoot", "Brown", "Smith", "CaddyWompus",
		"Cumberbatch", "Skeever-hole-tweed"}

	var nameList []string

	for i := 0; i < number; i++ {
		name := first[r.Intn(len(first))] + " " + last[r.Intn(len(last))]
		nameList = append(nameList, name)
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
		numList = append(numList, r.Int31n(1048500))
	}

	return numList
}
