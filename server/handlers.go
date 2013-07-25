package server

import (
	//"encoding/json"
	"errors"
	"fmt"
	"github.com/learc83/toastyserver/database"
	"log"
	"net/http"
	"strconv"
)

//http handlers--result should be returned as a hashmap with an
//"error" key, and a data key. Example: result["name"] = "jane",
//result["error"] = ""

func employeeLogin(req *http.Request, result map[string]interface{}) {
	keyNum := req.FormValue("KeyfobNumber")
	keyNumInt, err := strconv.Atoi(keyNum) //Atoi shortcut for ParseInt(s,10,0)
	if err != nil {
		result["error"] = stringifyErr(err, "employeeLogin()")
		return
	}

	name, err := database.FindEmployee(keyNumInt)
	if err != nil {
		result["error"] = stringifyErr(err, "employeeLogin()")
		return
	}

	result["name"] = name
}

func customerList(req *http.Request, result map[string]interface{}) {
	customers, err := database.RecentFiftyCustomers()
	if err != nil {
		result["error"] = stringifyErr(err, "customerList()")
		return
	}

	result["customers"] = customers
}

func customerListByName(req *http.Request, result map[string]interface{}) {
	name := req.FormValue("CustomerName")
	if name == "" {
		err := errors.New("CustomerName param blank")
		result["error"] = stringifyErr(err, "customerListByName()")
		return
	}

	customers, err := database.FindCustomersByName(name)
	if err != nil {
		result["error"] = stringifyErr(err, "customerListByName()")
		return
	}

	result["customers"] = customers
}

func addNewCustomer(req *http.Request, result map[string]interface{}) {
	name := req.FormValue("CustomerName")
	//name := req.FormValue("CustomerName")

	if name == "" {
		err := errors.New("CustomerName param blank")
		result["error"] = stringifyErr(err, "customerListByName()")
		return
	}
}

func availableKeyfobs(req *http.Request, result map[string]string) {

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
