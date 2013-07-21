package server

import (
	"github.com/learc83/toastyserver/database"
	"log"
	"net/http"
)

//wraps handlers to add default headers
func defaultHeaders(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		handler(w, r)
	}
}

func StartServer() {
	//WARNING -- DevMode DELETES DB
	database.OpenDBDevMode()

	routes := getRoutes()
	for key, value := range routes {
		http.HandleFunc(key, defaultHeaders(value))
	}

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
