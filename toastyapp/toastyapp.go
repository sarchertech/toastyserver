package main

import (
	"github.com/learc83/toastyserver/server"
	"runtime"
)

func main() {
	//Set max number of OS threads, no default so must set here
	runtime.GOMAXPROCS(1)

	server.StartServer()
}
