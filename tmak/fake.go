package tmak

import (
	"github.com/learc83/toastyserver/database"
	"time"
)

//Side Effects: edits beds in place
func BedStatuses(beds []database.Bed) (err error) {
	for i := range beds {
		beds[i].Status = !(i%3 == 0)
	}

	return
}

func StartBed(bed int) (err error) {
	time.Sleep(1.5 * 1e9)

	return
}
