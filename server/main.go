package main

import (
	"fmt"

	"github.com/justyntemme/itsaRap/server/restAPI"
)

func main() {
	fmt.Println("We work!")
	go func() {
		restAPI.Run()
	}()

}
