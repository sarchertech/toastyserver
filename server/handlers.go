package server

import (
	"encoding/json"
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
	eNum := req.FormValue("EmployeeNumber")

	found := true //FindEmployee(eNum)

	if found {
		w.Write([]byte("Authorized " + eNum))
	} else {
		w.Write([]byte("Not Authorized"))
	}
}
