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
	for i:=0; i<5; i++ {
		time.Sleep(4000 * time.Millisecond)
		newValue := lightValue + 50

		lightValue = newValue

		dmx.SetChannel(3, byte(lightValue))
		dmx.Render()
		fmt.Println(fmt.Sprintf("%s %d", "Channel 3 @", lightValue))
	}

	fmt.Println("Finished")
}

