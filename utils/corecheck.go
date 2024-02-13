package utils

import (
	"fmt"
	"log"
	"runtime"
)

func CoreCheck(cores, verbose *int) {
	if *cores > runtime.NumCPU() || *cores <= 0 {
		log.Fatal(`You don't have that many cores... you can use up to `,
			runtime.NumCPU())
	} else {
		runtime.GOMAXPROCS(*cores)
		if *verbose == 1 {
			fmt.Println("Using " + string(runtime.GOMAXPROCS(*cores)) + " cores")
		}
	}
}
