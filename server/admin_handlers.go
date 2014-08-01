package server

import (
	"github.com/learc83/toastyserver/database"
	"net/http"
)

//http handlers--result should be returned as a hashmap with an
//"error" key, and a data key. Example: result["name"] = "jane",
//result["error"] = ""

//TODO replace stringifyErr with fmt.ErrorF() now that I know it exists

func employeeLogin(req *http.Request, result map[string]interface{}) {
	params, err := getParams(req, param{"Fob_num", "int"})
	if err != nil {
		result["error"] = stringifyErr(err, "Error Logging In")
		return
	}

	name, err := database.FindEmployee(params["Fob_num"].(int))
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

//TODO enforce non-blank customer name string
func addNewCustomer(req *http.Request, result map[string]interface{}) {
	params, err := getParams(req,
		param{"name", "str"},
		param{"phone_number", "str"},
		param{"level", "int"},
		param{"keyfob_number", "int"})

	if err != nil {
		result["error"] = stringifyErr(err, "Error Adding New Customer")
		return
	}

	customer := database.Customer{
		Name:    params["name"].(string),
		Phone:   params["phone_number"].(string),
		Status:  true,
		Level:   params["level"].(int),
		Fob_num: params["keyfob_number"].(int)}

	err = database.CreateRecord(customer)

	if err != nil {
		result["error"] = stringifyErr(err, "Error Adding New Customer")
		return
	}
}

func deleteCustomer(req *http.Request, result map[string]interface{}) {
	params, err := getParams(req, param{"customer_id", "int"})

	if err != nil {
		result["error"] = stringifyErr(err, "Error Deleting Customer")
		return
	}

	//WARNING doesn't return error if record doesn't exist
	err = database.DeleteCustomer(params["customer_id"].(int))

	if err != nil {
		result["error"] = stringifyErr(err, "Error Deleting Customer")
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

func doorReport(req *http.Request, result map[string]interface{}) {
	accesses, err := database.RecentDoorAccesses() //500
	if err != nil {
		result["error"] = stringifyErr(err, "Error Displaying Door Report")
		return
	}

	result["doorAccesses"] = accesses
}

func tanReport(req *http.Request, result map[string]interface{}) {
	sessions, err := database.RecentTanSessions() //500
	if err != nil {
		result["error"] = stringifyErr(err, "Error Displaying Tan Report")
		return
	}

	result["tanSessions"] = sessions
}

//TODO enforce non-blank bed name string
func addNewBed(req *http.Request, result map[string]interface{}) {
	params, err := getParams(req,
		param{"level", "int"},
		param{"max_time", "int"},
		param{"name", "string"})

	if err != nil {
		result["error"] = stringifyErr(err, "Error Adding New Bed")
		return
	}

	bed := database.Bed{
		Level:	  params["level"].(int),
		Max_time: params["max_time"].(int),
		Name:	  params["name"].(string)}

	err = database.CreateRecord(bed)

	if err != nil {
		result["error"] = stringifyErr(err, "Error Adding New Bed")
		return
	}
}

func deleteBed(req *http.Request, result map[string]interface{}) {
	params, err := getParams(req, param{"bed_num", "int"})

	if err != nil {
		result["error"] = stringifyErr(err, "Error Deleting Bed")
		return
	}

	//WARNING doesn't return error if record doesn't exist
	err = database.DeleteBed(params["bed_num"].(int))

	if err != nil {
		result["error"] = stringifyErr(err, "Error Deleting Bed")
		return
	}
}

func updateBed(req *http.Request, result map[string]interface{}) {
	params, err := getParams(req,
		param{"bed_num", "int"},
		param{"level", "int"},
		param{"max_time", "int"},
		param{"name", "string"})

	if err != nil {
		result["error"] = stringifyErr(err, "Error Updating Bed")
		return
	}

	bed := database.Bed{
		Bed_num:  params["bed_num"].(int),
		Level:	  params["level"].(int),
		Max_time: params["max_time"].(int),
		Name:	  params["name"].(string)}

	err = database.UpdateBed(bed)

	if err != nil {
		result["error"] = stringifyErr(err, "Error Updating Bed")
		return
	}
}

func listBeds(req *http.Request, result map[string]interface{}) {
	beds, err := database.ListBeds()
	if err != nil {
		result["error"] = stringifyErr(err, "Error Displaying Bed List")
		return
	}

	result["beds"] = beds
}