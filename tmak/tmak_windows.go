// +build windows

package tmak

import (
	"fmt"
	"github.com/learc83/toastyserver/database"
	"time"
)

func init() {
	fmt.Println("Fake Tmax Started")
}

//Side Effects: edits beds in place
func BedStatuses(beds []database.Bed) (err error) {
	for i := range beds {
		beds[i].Status = !(beds[i].Bed_num%3 == 0)
	}

	fmt.Println("Fake Bed Statuses Called")

	return
}

func StartBed(bed int, t int) (err error) {
	time.Sleep(1.5 * 1e9)

	fmt.Println("Fake Bed Started")

	return
}
