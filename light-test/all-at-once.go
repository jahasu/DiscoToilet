package main

import (
	"github.com/akualab/dmx"
	"fmt"
	
)

func main () {
	dmx, e := dmx.NewDMXConnection("/dev/ttyUSB0")
	if e != nil {
		fmt.Println("ERROR: Unable to connect to DMX interface.")
		return
	}
	defer dmx.Close()

	fmt.Println("Start")

	dmx.SetChannel(3, 200)
	dmx.Render()
	fmt.Println("Channel 3 => 200")

	dmx.SetChannel(4, 200)
	dmx.Render()
	fmt.Println("Channel 3 & 4 => 200")

	dmx.SetChannel(5, 200)
	dmx.Render()
	fmt.Println("Channel 3, 4, 5 => 200")
}

