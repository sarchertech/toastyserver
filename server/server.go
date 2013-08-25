package server

import (
	"github.com/learc83/toastyserver/database"
	"log"
	"net/http"
)

func StartServer() {
	database.OpenDB()

	for key, value := range getRoutes() {
		http.HandleFunc(key, handlerWrapper(value))
	}

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
