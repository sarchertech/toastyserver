// +build !door

package door

import (
	"fmt"
	"time"
	"github.com/learc83/toastyserver/database"
)

func StartDoorControl() {
	fmt.Println("Door control not enabled.")

	for {
		fmt.Println("Door control not enabled.")
		customers, _ := database.RecentFiftyCustomers()
		fmt.Println(customers)

		time.Sleep(1.5 * 1e9)
	}	
}