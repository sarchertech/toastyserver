package server

import (
//"net/http"
)

//routes to match handlers to url strings

func getRoutes() map[string]toastyHndlrFnc {
	r := make(map[string]toastyHndlrFnc)

	//admin routes
	r["/employee_login"] = employeeLogin
	r["/customer_list"] = customerList
	r["/customer_list_by_name"] = customerListByName
	r["/add_new_customer"] = addNewCustomer
	r["/available_customer_keyfobs"] = availableCustomerKeyfobs

	//customer routes
	r["/customer_login"] = customerLogin
	r["/bed_status"] = bedStatus
	r["/start_bed"] = startBed

	return r
}
