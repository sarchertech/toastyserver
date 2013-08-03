package server

import (
	"github.com/learc83/toastyserver/database"
	"net/http"
)

//http handlers--result should be returned as a hashmap with an
//"error" key, and a data key. Example: result["name"] = "jane",
//result["error"] = ""

func customerLogin(req *http.Request, result map[string]interface{}) {
	params, err := getParams(req, param{"Fob_num", "int"})
	if err != nil {
		result["error"] = stringifyErr(err, "Error Logging In")
		return
	}

	name, stat, lvl, err := database.FindCustomer(params["Fob_num"].(int))
	if err != nil {
		result["error"] = stringifyErr(err, "Error Logging In")
		return

	}
	result["name"] = name
	result["status"] = stat
	result["level"] = lvl
}
