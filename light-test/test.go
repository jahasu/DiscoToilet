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

	fmt.Println("Start")

	lightValue := 0
	dmx.SetChannel(2, byte(200))
	for i:=0; i<5; i++ {
		time.Sleep(4000 * time.Millisecond)
		newValue := lightValue + 50

		lightValue = newValue

		dmx.SetChannel(6, byte(lightValue))
		dmx.Render()
		fmt.Println(fmt.Sprintf("%s %d", "Channel 6 @", lightValue))
	}

	fmt.Println("Finished")
}

