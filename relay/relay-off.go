package main

import (
	"fmt"

	"github.com/akualab/dmx"
	"github.com/stianeikeland/go-rpio"
)

func resetLights() {
	dmx, e := dmx.NewDMXConnection("/dev/ttyUSB0")
	if e != nil {
		fmt.Println("ERROR: Unable to connect to DMX interface.")
		return
	}
	defer dmx.Close()

	// fmt.Println("Start")

	for i := 0; i <= 100; i++ {
		dmx.SetChannel(i+1, 0)
		dmx.Render()
		//fmt.Println(fmt.Sprintf("%s %d %s", "Channel ", i, "set to 0"))
	}

	fmt.Println("DMX channels reset")
	dmx.Close()
}

func main() {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		fmt.Println(fmt.Sprint("unable to open gpio", err.Error()))
	}

	relay1 := rpio.Pin(23)
	relay2 := rpio.Pin(24)

	relay1.Output()
	relay2.Output()

	relay1.Low()
	relay2.Low()

	resetLights()
}
