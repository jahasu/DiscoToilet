package main

import (
	"github.com/akualab/dmx"
	"fmt"
	"time"
)

func main () {
	dmx, e := dmx.NewDMXConnection("/dev/ttyUSB0")
	if e != nil {
		fmt.Println("ERROR: Unable to connect to DMX interface.")
		return
	}
	defer dmx.Close()

	dmx.SetChannel(1, 200)
	dmx.Render()
	fmt.Println("Channel 1 => 200")
	time.Sleep(5000 * time.Millisecond)

	dmx.SetChannel(1, 0)
	dmx.SetChannel(2, 200)
	dmx.Render()
	fmt.Println("Channel 2 => 200")
	time.Sleep(5000 * time.Millisecond)

        dmx.SetChannel(2, 0)
	dmx.SetChannel(3, 200)
	dmx.Render()
	fmt.Println("Channel 3 => 200")
	time.Sleep(5000 * time.Millisecond)

        dmx.SetChannel(3, 0)
	dmx.SetChannel(4, 200)
	dmx.Render()
	fmt.Println("Channel 4 => 200")
	time.Sleep(5000 * time.Millisecond)

        dmx.SetChannel(4, 0)
	dmx.SetChannel(5, 200)
	dmx.Render()
	fmt.Println("Channel 5 => 200")
	time.Sleep(5000 * time.Millisecond)

        dmx.SetChannel(5, 0)
	dmx.SetChannel(6, 200)
	dmx.Render()
	fmt.Println("Channel 6 => 200")
	time.Sleep(5000 * time.Millisecond)
}
