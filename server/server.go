package server

import (
	"encoding/json"
	"github.com/learc83/toastyserver/database"
	"log"
	"net/http"
)

type toastyHndlrFnc func(*http.Request, *map[string]interface{})

func handlerWrapper(handler toastyHndlrFnc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		//result gets set inside the handler function, would have liked
		//to keep this more functional, but this is more efficient for
		//large return data
		result := make(map[string]interface{})
		handler(r, &result)

		j, err := json.Marshal(result)
		if err != nil {
			log.Println(err)
			errs := `{"error": "json.Marshal failed", "name": ""}`
			w.Write([]byte(errs))
			return
		}
		w.Write(j)
	}
}

func StartServer() {
	//WARNING -- DevMode DELETES DB
	database.OpenDBDevMode()

	routes := getRoutes()
	for key, value := range routes {
		http.HandleFunc(key, handlerWrapper(value))
	}

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
