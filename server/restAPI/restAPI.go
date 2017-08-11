package restAPI

import (
	"fmt"
)

//Run Starts Go REST API server
func Run() {

	go func() {
		fmt.Print("goRoutine Started!")
	}()
}
