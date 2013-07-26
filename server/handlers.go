package server

import (
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
	params, err := getParams(req, param{"keyNum", "int"})
	if err != nil {
		result["error"] = stringifyErr(err, "Error Logging In")
		return
	}

	name, err := database.FindEmployee(params["keyNum"].(int))
	if err != nil {
		result["error"] = stringifyErr(err, "Error Logging In")
		return
	}

	result["name"] = name
}

func customerList(req *http.Request, result map[string]interface{}) {
	customers, err := database.RecentFiftyCustomers()
	if err != nil {
		result["error"] = stringifyErr(err, "Error Displaying Customer List")
		return
	}

	result["customers"] = customers
}

func customerListByName(req *http.Request, result map[string]interface{}) {
	params, err := getParams(req, param{"name", "str"})
	if err != nil {
		result["error"] = stringifyErr(err, "Error Searching Customers")
		return
	}

	customers, err := database.FindCustomersByName(params["name"].(string))
	if err != nil {
		result["error"] = stringifyErr(err, "Error Searching Customers")
		return
	}

	result["customers"] = customers
}

func addNewCustomer(req *http.Request, result map[string]interface{}) {
	params, err := getParams(req,
		param{"name", "str"},
		param{"phone number", "str"},
		param{"level", "int"},
		param{"keyfob number", "int"})

	if err != nil {
		result["error"] = stringifyErr(err, "Error Adding New Customer")
		return
	}

	err = database.CreateCustomer(
		params["name"].(string),
		params["phone number"].(string),
		params["level"].(int),
		params["keyfob number"].(int))

	if err != nil {
		result["error"] = stringifyErr(err, "Error Adding New Customer")
		return
	}
}

func availableCustomerKeyfobs(req *http.Request, result map[string]interface{}) {
	keyfobsTen, keyfobsHex, err := database.AvailableCustomerKeyfobs()
	if err != nil {
		result["error"] = stringifyErr(err, "Error Finding Available Customer Keyfobs")
		return
	}

	result["keyfobsTen"] = keyfobsTen
	result["keyfobsHex"] = keyfobsHex
}

//used for get Params arguments. Only supports ints and strings, add support for
//checking things like phone numbers
type param struct {
	Name string
	Type string
}

func getParams(req *http.Request, paramList ...param) (params map[string]interface{}, err error) {
	params = make(map[string]interface{})
	blanks := ""
	notInts := ""

	for _, p := range paramList {
		param := req.FormValue(p.Name)
		if param == "" {
			blanks = blanks + " " + p.Name + ","
			continue
		}

		if p.Type == "int" {
			num, errr := strconv.Atoi(param) //Atoi shortcut for ParseInt(s,10,0)
			if errr != nil {
				notInts = notInts + " " + p.Name + ","
				continue
			}
			params[p.Name] = num
		} else {
			params[p.Name] = param
		}
	}

	if blanks != "" {
		err = errors.New("These fields cannot be blank:" + blanks)
	}

	if notInts != "" { //TODO modify so it prints blank for nil err instead of <nil>
		err = errors.New(fmt.Sprintf("%v. These fields must be numbers:%s", err, notInts))
	}

	return
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
