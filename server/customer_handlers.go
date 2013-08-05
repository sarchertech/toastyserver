package server

import (
	"github.com/learc83/toastyserver/database"
	"github.com/learc83/toastyserver/tmak"
	"log"
	//"github.com/learc83/toastyserver/tmak"
	"net/http"
)

//http handlers--result should be returned as a hashmap with an
//"error" key, and a data key. Example: result["name"] = "jane",
//result["error"] = ""

func customerLogin(req *http.Request, result map[string]interface{}) {
	params, err := getParams(req, param{"Fob_num", "int"})
	if err != nil {
		result["error"] = stringifyErr(err, "Error With Customer Login")
		return
	}

	name, stat, lvl, err := database.FindCustomer(params["Fob_num"].(int))
	if err != nil {
		result["error"] = stringifyErr(err, "Error With Customer Login")
		return

	}
	result["name"] = name
	result["status"] = stat
	result["level"] = lvl
}

func bedStatus(req *http.Request, result map[string]interface{}) {
	params, err := getParams(req, param{"Level", "int"})
	if err != nil {
		result["error"] = stringifyErr(err, "Error Checking Customer Bed Status")
		return
	}

	beds, err := database.BedsByLevel(params["Level"].(int))
	if err != nil {
		result["error"] = stringifyErr(err, "rror Checking Customer Bed Status")
		return
	}
	log.Println(beds)

	tmak.BedStatuses(beds)

	result["beds"] = beds
}
