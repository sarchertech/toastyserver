// +build linux
// +build production

//use go install -tags production

package tmak

import (
	"errors"
	"github.com/learc83/toastyserver/database"
	"github.com/schleibinger/sio"
	"log"
	"sync"
	"syscall"
	"time"
)

type tmak struct {
	mu   sync.Mutex
	port *sio.Port
}

var T tmak

func init() {
	var err error
	T.port, err = sio.Open("/dev/ttyUSB0", syscall.B115200)
	if err != nil {
		log.Println(err)
	}

	log.Println(T.port)

	//StartBed(7)
}

//Side Effects: edits beds in place
func BedStatuses(beds []database.Bed) (err error) {
	rBuf := make([]byte, 37) //37 bytes--3 start, 1 command, 32 data, 1 chksum
	err = tryBedStatus(rBuf)
	if err != nil {
		T.port.Close()
		T.port, err = sio.Open("/dev/ttyUSB0", syscall.B115200)
		time.Sleep(0.020 * 1e9)
		err = tryBedStatus(rBuf)
		if err != nil {
			T.port.Close()
			T.port, err = sio.Open("/dev/ttyUSB0", syscall.B115200)
			time.Sleep(0.020 * 1e9)
			err = tryBedStatus(rBuf)
		}
	}

	for i := range beds {
		s := rBuf[beds[i].Bed_num+3]

		beds[i].Status = (s == 0 || s == 4)
	}

	return
}

func tryBedStatus(buf []byte) (err error) {
	//TODO WARNING write error handling for uint8 conversion
	//bytes 4 and 5 are start and end for # of beds returned
	n, err := T.port.Write([]byte{255, 254, 253, 4, 1, 32, 0, 0})
	if err != nil {
		log.Println(err)
		return
	}

	time.Sleep(0.020 * 1e9) //5ms

	n, err = T.port.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}
	if n < 37 {
		log.Println("Short Read in Bed Status")
		err = errors.New("Short Read Error in Bed Status")
		log.Println(buf)
		return
	}
	if !startBytesCorrect(buf) {
		log.Println("Starting Bytes not correct in Bed Status")
		err = errors.New("Starting Bytes bad Error in Bed Status")
		log.Println(buf)
		return
	}
	if !chksumCorrectStatus(buf) {
		log.Println("Chksum bad in Bed Status")
		err = errors.New("Chksum Error in Bed Status")
		log.Println(buf)
		return
	}

	log.Println(n)
	log.Println(buf)
	return
}

//TODO WARNING this is AWFUL--FIX IT. The port closes and to resync the stream,
//b/c the serial library we are using has no flush method exposed.
func StartBed(bed int, time int) (err error) {
	err = tryStartBed(bed, time)
	if err != nil {
		T.port.Close()
		T.port, err = sio.Open("/dev/ttyUSB0", syscall.B115200)
		err = tryStartBed(bed, time)
		if err != nil {
			T.port.Close()
			T.port, err = sio.Open("/dev/ttyUSB0", syscall.B115200)
			err = tryStartBed(bed, time)
		}
	}
	return
}

func tryStartBed(bed int, t int) (err error) {
	//TODO WARNING write error handling for uint8 conversion
	n, err := T.port.Write([]byte{255, 254, 253, 1, uint8(bed), uint8(t), 5, 0})
	if err != nil {
		log.Println(err)
		return
	}

	//TODO WARNING delay may need to be changed
	time.Sleep(0.005 * 1e9) //5ms

	rxbuf := make([]byte, 8)

	n, err = T.port.Read(rxbuf)
	if err != nil {
		log.Println(err)
		return
	}
	if n < 8 {
		log.Println("Short Read in Bed Start")
		err = errors.New("Short Read Error in Bed Start")
		return
	}
	if !startBytesCorrect(rxbuf) {
		log.Println("Starting Bytes not correct in Bed Start")
		err = errors.New("Starting Bytes bad Error in Bed Start")
		log.Println(rxbuf)
		return
	}
	if !chksumCorrect(rxbuf) {
		log.Println("Chksum bad in Bed Start")
		err = errors.New("Chksum Error in Bed Start")
		log.Println(rxbuf)
		return
	}

	//log.Println(n)
	log.Println(rxbuf)

	return
}

func startBytesCorrect(buf []byte) (correct bool) {
	if buf[0] == 255 && buf[1] == 254 && buf[2] == 253 {
		correct = true
	}

	return
}

func chksumCorrect(chksum []byte) (correct bool) {
	var sum int

	for i := 0; i < 7; i++ {
		sum = sum + int(chksum[i])
		//log.Println(i)
		//log.Println(sum)
	}

	if chksum[7] == uint8(sum%255) {
		correct = true
	}

	log.Println(sum)

	return
}

func chksumCorrectStatus(chksum []byte) (correct bool) {
	var sum int

	for i := 0; i < 36; i++ {
		sum = sum + int(chksum[i])
		//log.Println(i)
		//log.Println(sum)
	}

	if chksum[36] == uint8(sum%255) {
		correct = true
	}

	log.Println(sum)

	return
}
