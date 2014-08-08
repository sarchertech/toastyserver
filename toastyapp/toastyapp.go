package main

import (
	"github.com/learc83/toastyserver/server"
	"runtime"
	"github.com/learc83/toastyserver/door"
)

func main() {
	//Set max number of OS threads, no default so must set here
	runtime.GOMAXPROCS(1)

	go door.StartDoorControl() //start door control in another thread

	server.StartServer()
}
