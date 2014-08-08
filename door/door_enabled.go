// +build door

//use go install -tags 'door other tags'

package door

import (
	"log"
	"github.com/learc83/sio"
	"syscall"
	"strconv"
	"github.com/learc83/toastyserver/database"
	"time"
)

//TODO remove most of the logs
func StartDoorControl() {
	log.Println("Door control enabled.")

	port, err := sio.Open("/dev/ttyUSB1", syscall.B9600)
	if err != nil {
		log.Println(err)
		return
	}

	// main loop
	for { 
		rBuf := make([]byte, 10) // tab, 8 byte hex keyfob #, carriage return


		// Read first byte until we get to a tab
		for {
			log.Println("reading first byte")
			_, err := port.Read(rBuf[:1]) 
			if err != nil {
				log.Println("error in first byte:")
				log.Println(err)
				port.Close()
				port, err = sio.Open("/dev/ttyUSB1", syscall.B9600)
				continue	
			}
			
			if rBuf[0] == 9 {
				break
			}

			log.Println("First byte wrong: ")
			log.Println(rBuf[0])
		}
		
		// Got first byte now read next 9 
		i := 1
		for i < 10 {	
			n, err := port.Read(rBuf[i:]) //make a slice so it reads up to 8	
			if err != nil {
				log.Println(err)
				break
			}
			i = i + n
		}

		if i < 10 {
			log.Println("Short read in door control")
			port.Close()
			port, err = sio.Open("/dev/ttyUSB1", syscall.B9600)
			if err != nil {
				log.Println(err)
				continue
			}
			continue
		}
		if (rBuf[0] != 9 && rBuf[9] != 13) {
			log.Println("Misformed message in door control")
			port.Close()
			port, err = sio.Open("/dev/ttyUSB1", syscall.B9600)
			if err != nil {
				log.Println(err)
				continue
			}
			continue
		}		

		i = 0
		for i < 10 {
			log.Println(rBuf[i])
			i++
		}

		s := string(rBuf[1:9])

		fobNum, _ := strconv.ParseUint(s, 16, 64)

		log.Println(s)
		log.Println(fobNum)

		id, _, _, _, err := database.FindCustomer(fobNum)
		if err != nil {
			log.Println(err)
			continue
		}

		//TODO log customers who aren't in the database
		//also add a check for status so that customers can
		//have tan access without door access 
		//default id value 0 if no customer found
		if id == 0 {
			log.Println("Door Access: keyfob not found")
			port.Write([]byte{9, 0, 0, 0, 13})
		} else {
			log.Println("Access granted")

			port.Write([]byte{9, 255, 254, 253, 13})

			doorAccess := database.DoorAccess{ Customer_id: id, 
				Time_stamp: time.Now().Unix()}
			
			database.CreateRecord(doorAccess)
		}
	}
}