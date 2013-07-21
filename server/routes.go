package server

import (
	"net/http"
)

//routes to match handlers to url strings

func getRoutes() map[string]http.HandlerFunc {
	r := make(map[string]http.HandlerFunc)

	//r["/json"] = helloServer
	r["/employee_login"] = employeeLogin

	return r
}
