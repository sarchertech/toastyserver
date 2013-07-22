package server

import (
//"net/http"
)

//routes to match handlers to url strings

func getRoutes() map[string]toastyHndlrFnc {
	r := make(map[string]toastyHndlrFnc)

	//r["/json"] = helloServer
	r["/employee_login"] = employeeLogin
	r["/customer_list"] = customerList
	r["/customer_list_by_name"] = customerListByName
	r["/customer_details"] = customerDetails
	r["/add_new_customer"] = addNewCustomer
	r["/available_keyfobs"] = availableKeyfobs

	return r
}
