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

	for i := 1; i <= 400; i++ {
		dmx.SetChannel(i+1, 0)
		dmx.Render()
		fmt.Println(fmt.Sprintf("%s %d %s", "Channel ", i, "set to 0"))
	}

	fmt.Println("Finished")
}
