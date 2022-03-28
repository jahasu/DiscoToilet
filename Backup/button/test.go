package main

import (
	"fmt"
	"os/exec"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	goExecPath, pathErr := exec.LookPath("go")

	if pathErr != nil {
		fmt.Println("Error: ", pathErr)
	} else {
		fmt.Println("Go Exectuable: ", goExecPath)
	}

	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		fmt.Println(fmt.Sprint("unable to open gpio", err.Error()))
	}

	defer rpio.Close()
	pin := rpio.Pin(3)

	pin.Input()
	pin.PullUp()

	for {
		fmt.Println(pin.Read())
	}
}
