package server

import (
	//"encoding/json"
	"github.com/learc83/toastyserver/database"
	"log"
	//"fmt"
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

//http handlers--result should be returned as a hashmap with an
//"error" key, and a data key. Example: result["name"] = "jane",
//result["error"] = ""

func employeeLogin(req *http.Request, result *map[string]string) {
	keyNum := req.FormValue("KeyfobNumber")

	rawResult, err := database.FindEmployee(keyNum)

	(*result)["name"] = rawResult
	(*result)["error"] = stringifyErr(err)
}

func customerList(req *http.Request, result *map[string]string) {
	//(*result)["customers"], (*result)["error"] = database.RecentFiftyCustomers()
}

func customerListByName(req *http.Request, result *map[string]string) {

}

func customerDetails(req *http.Request, result *map[string]string) {

}

func addNewCustomer(req *http.Request, result *map[string]string) {

}

func availableKeyfobs(req *http.Request, result *map[string]string) {

}

//This is required because errors default strinfigy method: Error()
//returns nil instead of an empty string
func stringifyErr(err error) string {
	if err != nil {
		log.Println(err)
		errs := err.Error()
		return errs
	}
	return ""
}
