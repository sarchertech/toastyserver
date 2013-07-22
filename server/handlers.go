package server

import (
	//"encoding/json"
	"github.com/learc83/toastyserver/database"
	//"log"
	"net/http"
)

//http handlers

// func helloServer(w http.ResponseWriter, req *http.Request) {
// 	type Message struct {
// 		Name           string
// 		CustomerNumber string
// 		Level          int8
// 		Message        string
// 	}

// 	qString := req.FormValue("CustomerNumber")

// 	var m Message

// 	if qString == "120121" {
// 		m = Message{"Bob Tanner", "120121", 3, ""}
// 	} else {
// 		m = Message{"Unkown", "", 0, "Can't find customer with that number"}
// 	}

// 	b, _ := json.Marshal(m)

// 	//w.Header().Set("Content-Type", "application/json")

// 	//qString := req.FormValue("butt")

// 	//io.WriteString(w, string(b))
// 	w.Write(b)
// }

func employeeLogin(w http.ResponseWriter, req *http.Request) {
	keyNum := req.FormValue("KeyfobNumber")

	result := make(map[string]string)

	result["name"], result["error"] = database.FindEmployee(keyNum)

	writeJson(w, result)
}

func customerList(w http.ResponseWriter, req *http.Request) {

}

func customerListByName(w http.ResponseWriter, req *http.Request) {

}

func customerDetails(w http.ResponseWriter, req *http.Request) {

}

func addNewCustomer(w http.ResponseWriter, req *http.Request) {

}

func availableKeyfobs(w http.ResponseWriter, req *http.Request) {

}
