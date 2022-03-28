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

	//Sphere
	dmx.SetChannel(1, 255) //dimmer
	dmx.SetChannel(2, 0)   //strobe
	dmx.SetChannel(3, 255) //red
	dmx.SetChannel(4, 255) //green
	dmx.SetChannel(5, 255) //blue
	dmx.SetChannel(6, 255) //rotate

	//Cage
	dmx.SetChannel(7, 255)  //dimmer
	dmx.SetChannel(8, 0)    //strobe
	dmx.SetChannel(9, 255)  //red
	dmx.SetChannel(10, 255) //green
	dmx.SetChannel(11, 255) //blue
	dmx.SetChannel(12, 255) //white
	dmx.SetChannel(13, 255) //rotate

	//Middle Light
	dmx.SetChannel(14, 10) //dimmer
	dmx.SetChannel(15, 0)  //strobe
	dmx.SetChannel(16, 30) //red
	dmx.SetChannel(17, 30) //green
	dmx.SetChannel(18, 30) //blue
	dmx.SetChannel(19, 30) //yellow

	//Laser
	dmx.SetChannel(20, 125) //laser r
	dmx.SetChannel(21, 125) //laser g
	dmx.SetChannel(22, 140) //animation & rotaion speed
	//dmx.SetChannel(23, 255)//???????

	dmx.Render()
	fmt.Println("Finished")
}
