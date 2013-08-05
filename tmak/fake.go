package tmak

import (
	"github.com/learc83/toastyserver/database"
)

//Side Effects: edits beds in place
func BedStatuses(beds []database.Bed) (err error) {
	for i := range beds {
		beds[i].Status = true
	}

	return
}
