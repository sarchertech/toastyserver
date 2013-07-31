package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type toastyHndlrFnc func(*http.Request, map[string]interface{})

func handlerWrapper(handler toastyHndlrFnc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result := make(map[string]interface{})
		handler(r, result) //result set as a side effect

		j, err := json.Marshal(result)
		if err != nil {
			log.Println(err)
			errs := `{"error": "json.Marshal failed"}`
			w.Write([]byte(errs))
			return
		}
		w.Write(j)
	}
}

//used for get Params arguments. Only supports ints and strings, add support for
//checking things like phone numbers
type param struct {
	Name string
	Type string
}

//TODO redo this function to take advantage of structs defined in database/structs.go
//and to use reflection. See CreateRecord function in database/sql.go
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
