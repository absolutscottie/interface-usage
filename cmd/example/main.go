package main

import (
	"fmt"

	"github.com/absolutscottie/interface-usage/container"
)

func main() {
	err := container.Initialize()
	if err != nil {
		//there was an error
		fmt.Printf("error during Initialize(): %v", err)
		return
	}

	err = container.UserDAO.Create()
	if err != nil {
		fmt.Printf("this should never happen")
	}
}
