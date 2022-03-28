package main

import (
	"fmt"

	"github.com/akualab/dmx"
)

func main() {
	dmx, e := dmx.NewDMXConnection("/dev/ttyUSB0")
	if e != nil {
		fmt.Println("ERROR: Unable to connect to DMX interface.")
		return
	}
	defer dmx.Close()

	fmt.Println("Start")

	dmx.SetChannel(1, 0)
	dmx.SetChannel(2, 0)
	dmx.SetChannel(3, 0)
	dmx.SetChannel(4, 0)
	dmx.SetChannel(5, 0)
	dmx.SetChannel(6, 0)
	dmx.SetChannel(7, 0)
	dmx.SetChannel(8, 0)
	dmx.SetChannel(9, 0)
	dmx.SetChannel(10, 0)

	dmx.Render()

	fmt.Println("Finished")
}
