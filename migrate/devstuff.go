package main

import (
	"fmt"
	"github.com/learc83/toastyserver/database"
	"math/rand"
)

const randSeed int64 = 56355145

var r *rand.Rand

func createDevelopmentDB() {
	database.DeleteDB()
	database.OpenDB()
	defer database.CloseDB()

	database.UpSchema()
	addDevData()

	// keyfob := Keyfob{Fob_num: 12107728, Admin: true}
	// database.CreateRecord(keyfob)

	// employee := Employee{Name: "Seth", Level: 3, Fob_num: 12107728}
	// database.CreateRecord(employee)

	// keyfob2 := Keyfob{Fob_num: 9873, Admin: false}
	// database.CreateRecord(keyfob2)

	// customer := Customer{Name: "Jane Tanner", Level: 3, Fob_num: 9873,
	// 	Phone: "770-949-1622", Status: true}
	// database.CreateRecord(customer)

	// customer2 := Customer{Name: "Fred Tanner", Level: 3, Fob_num: 9871,
	// 	Phone: "770-949-1622", Status: false}
	// database.CreateRecord(customer2)

	// session := Session{
	// 	Bed_num:     5,
	// 	Customer_id: 11,
	// 	Time_stamp:  time.Now().Unix() - 43201}

	// err := database.CreateRecord(session)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// session.Time_stamp = time.Now().Unix() - 50000

	// err = database.CreateRecord(session)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

}

func addDevData() {
	//initialize random seed
	r = rand.New(rand.NewSource(randSeed))

	adminKeyfobs, customerKeyfobs := addFakeKeyfobs()
	addFakeEmployees(adminKeyfobs)
	addFakeCustomers(customerKeyfobs)
	addFakeBeds()
}

func addFakeKeyfobs() (adminKeyfobs []int, customerKeyfobs []int) {
	adminKeyfobs = fakeNumbers(10)
	customerKeyfobs = fakeNumbers(10)

	for k := range adminKeyfobs {
		keyfob := database.Keyfob{Fob_num: adminKeyfobs[k], Admin: true}
		database.CreateRecord(keyfob)
	}

	for k := range customerKeyfobs {
		keyfob := database.Keyfob{Fob_num: customerKeyfobs[k], Admin: false}
		database.CreateRecord(keyfob)
	}

	return
}

func addFakeEmployees(keyfobs []int) {
	for e := range keyfobs {
		employee := database.Employee{Name: fakeName(), Level: 1, Fob_num: keyfobs[e]}
		database.CreateRecord(employee)
	}
}

func addFakeCustomers(keyfobs []int) {
	for e := range keyfobs {
		customer := database.Customer{Name: fakeName(), Level: 3, Fob_num: keyfobs[e],
			Phone: fakePhone(), Status: true}
		database.CreateRecord(customer)
	}
}

func addFakeBeds() {
	//Need to fix numbering, quick hax
	for i := 0; i < 10; i++ {
		bed := database.Bed{Bed_num: i + 1, Level: 1, Max_time: 15, Name: "Sundash 232"}
		database.CreateRecord(bed)
	}

	for i := 0; i < 5; i++ {
		bed := database.Bed{Bed_num: i + 11, Level: 2, Max_time: 12, Name: "Ameribed 64"}
		database.CreateRecord(bed)
	}

	for i := 0; i < 2; i++ {
		bed := database.Bed{Bed_num: i + 16, Level: 3, Max_time: 10, Name: "Bad Ass Bed"}
		database.CreateRecord(bed)
	}
}

func fakeName() string {
	first := []string{"Bob", "Susanne", "Jennifer", "Georginamar", "Betty",
		"Grant", "Sarah", "Loranne", "Zorahflordian", "Seven"}

	last := []string{"Robertson", "Franklin-Louis", "Mc Van Derson-Eberts",
		"Pickles", "Wintergreen", "von Snoot", "Brown", "Smith", "CaddyWompus",
		"Cumberbatch", "Skeever-hole-tweed"}

	return first[r.Intn(len(first))] + " " + last[r.Intn(len(last))]
}

func fakePhone() string {
	area := []string{"770", "404", "680", "755", "804", "925"}

	first := fmt.Sprintf("%d%d%d", r.Intn(10), r.Intn(10), r.Intn(10))
	last := fmt.Sprintf("%d%d%d%d", r.Intn(10), r.Intn(10), r.Intn(10), r.Intn(10))

	return area[r.Intn(len(area))] + "-" + first + "-" + last
}

//could make duplicates
func fakeNumbers(number int) []int {
	var numList []int

	for i := 0; i < number; i++ {
		numList = append(numList, int(r.Int31n(1048500)))
	}

	return numList
}
