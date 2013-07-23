package server

import (
	//"encoding/json"
	"fmt"
	"github.com/learc83/toastyserver/database"
	"log"
	"net/http"
	"strconv"
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

func employeeLogin(req *http.Request, result *map[string]interface{}) {
	keyNum := req.FormValue("KeyfobNumber")

	keyNumInt, err := strconv.Atoi(keyNum) //Atoi shortcut for ParseInt(s,10,0)
	if err != nil {
		(*result)["error"] = stringifyErr(err, "employeeLogin()")
		return
	}

	data, err := database.FindEmployee(keyNumInt)

	(*result)["name"] = data
	(*result)["error"] = stringifyErr(err, "employeeLogin()")
}

func customerList(req *http.Request, result *map[string]interface{}) {
	//(*result)["customers"], (*result)["error"] = database.RecentFiftyCustomers()

	data, err := database.RecentFiftyCustomers()

	(*result)["customers"] = data
	(*result)["error"] = stringifyErr(err, "customerList()")
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
func stringifyErr(err error, callingFunc string) string {
	if err != nil {
		errs := fmt.Sprintf("%s: %s", callingFunc, err)
		log.Println(errs)
		return errs
	}
	return ""
}
