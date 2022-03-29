package main

import (
	"fmt"
	"os/exec"
	"time"

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

func turnOnLights() {
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
	dmx.SetChannel(14, 255) //dimmer
	dmx.SetChannel(15, 0)   //strobe
	dmx.SetChannel(16, 255) //red
	dmx.SetChannel(17, 255) //green
	dmx.SetChannel(18, 255) //blue
	dmx.SetChannel(19, 255) //yellow

	//Laser
	dmx.SetChannel(20, 255) //laser r
	dmx.SetChannel(21, 255) //laser g
	dmx.SetChannel(22, 255) //animation & rotaion speed
	dmx.SetChannel(23, 255) //???????
	//-------------------------------------------
	dmx.SetChannel(30, 255) //dimmer   ball in the middle
	dmx.SetChannel(31, 0)   //strobe
	dmx.SetChannel(32, 255) //red  brokem
	dmx.SetChannel(33, 0)   //green
	dmx.SetChannel(34, 255) //blue  beoken
	dmx.SetChannel(35, 255) //moon rotate
	dmx.SetChannel(36, 255) //dimmer
	dmx.SetChannel(37, 0)   //strobe
	dmx.SetChannel(38, 255) //red
	dmx.SetChannel(39, 0)   //green
	dmx.SetChannel(40, 255) //blue
	dmx.SetChannel(41, 0)   //yellow
	dmx.SetChannel(42, 255) //laser r
	dmx.SetChannel(43, 0)   //laser g
	dmx.SetChannel(44, 255) //pan rotate
	//-------------------------------------------
	dmx.SetChannel(60, 255) //dimmer ball on toilet
	dmx.SetChannel(61, 0)   //strobe
	dmx.SetChannel(62, 255) //red (Broken)
	dmx.SetChannel(63, 0)   //green (Broken)
	dmx.SetChannel(64, 255) //blue
	dmx.SetChannel(65, 255) //moon rotate
	dmx.SetChannel(66, 255) //dimmer
	dmx.SetChannel(67, 0)   //strobe  bad
	dmx.SetChannel(68, 255) //red
	dmx.SetChannel(69, 0)   //green  stuck on
	dmx.SetChannel(70, 255) //blue   stuck
	dmx.SetChannel(71, 0)   //yellow  stuck
	dmx.SetChannel(72, 0)   //laser r
	dmx.SetChannel(73, 0)   //laser g
	dmx.SetChannel(74, 0)   //pan rotate

	dmx.Render()
	fmt.Println("Finished")
	dmx.Close()
}

func turnOnLightsBossMode() {
	dmx, e := dmx.NewDMXConnection("/dev/ttyUSB0")
	if e != nil {
		fmt.Println("ERROR: Unable to connect to DMX interface.")
		return
	}
	defer dmx.Close()

	fmt.Println("Start")

	//Sphere
	dmx.SetChannel(1, 255) //dimmer
	dmx.SetChannel(2, 30)  //strobe
	dmx.SetChannel(3, 255) //red
	dmx.SetChannel(4, 255) //green
	dmx.SetChannel(5, 255) //blue
	dmx.SetChannel(6, 255) //rotate

	//Cage
	dmx.SetChannel(7, 255)  //dimmer
	dmx.SetChannel(8, 30)   //strobe
	dmx.SetChannel(9, 255)  //red
	dmx.SetChannel(10, 255) //green
	dmx.SetChannel(11, 255) //blue
	dmx.SetChannel(12, 255) //white
	dmx.SetChannel(13, 255) //rotate

	//Middle Light
	dmx.SetChannel(14, 255) //dimmer
	dmx.SetChannel(15, 30)  //strobe
	dmx.SetChannel(16, 255) //red
	dmx.SetChannel(17, 255) //green
	dmx.SetChannel(18, 255) //blue
	dmx.SetChannel(19, 255) //yellow

	//Laser
	dmx.SetChannel(20, 255) //laser r
	dmx.SetChannel(21, 255) //laser g
	dmx.SetChannel(22, 255) //animation & rotaion speed
	dmx.SetChannel(23, 255) //???????
	//-------------------------------------------
	dmx.SetChannel(30, 255) //dimmer   ball in the middle
	dmx.SetChannel(31, 255) //strobe
	dmx.SetChannel(32, 255) //red  brokem
	dmx.SetChannel(33, 200) //green
	dmx.SetChannel(34, 255) //blue  beoken
	dmx.SetChannel(35, 255) //moon rotate
	dmx.SetChannel(36, 255) //dimmer
	dmx.SetChannel(37, 0)   //strobe
	dmx.SetChannel(38, 255) //red
	dmx.SetChannel(39, 255) //green
	dmx.SetChannel(40, 255) //blue
	dmx.SetChannel(41, 255) //yellow
	dmx.SetChannel(42, 255) //laser r
	dmx.SetChannel(43, 20)  //laser g
	dmx.SetChannel(44, 255) //pan rotate
	//-------------------------------------------
	dmx.SetChannel(60, 255) //dimmer ball on toilet
	dmx.SetChannel(61, 0)   //strobe
	dmx.SetChannel(62, 255) //red (Broken)
	dmx.SetChannel(63, 0)   //green (Broken)
	dmx.SetChannel(64, 255) //blue
	dmx.SetChannel(65, 255) //moon rotate
	dmx.SetChannel(66, 255) //dimmer
	dmx.SetChannel(67, 0)   //strobe  bad
	dmx.SetChannel(68, 255) //red
	dmx.SetChannel(69, 0)   //green  stuck on
	dmx.SetChannel(70, 255) //blue   stuck
	dmx.SetChannel(71, 50)  //yellow  stuck
	dmx.SetChannel(72, 50)  //laser r
	dmx.SetChannel(73, 20)  //laser g
	dmx.SetChannel(74, 200) //pan rotate

	dmx.Render()
	fmt.Println("Finished")
	dmx.Close()
}

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
	relay1 := rpio.Pin(23)
	relay2 := rpio.Pin(24)

	power := 0
	lastPinRead := 0
	pressCount := 0
	count := 0
	pinState := 0

	pin.Input()
	pin.PullUp()

	relay1.Output()
	relay2.Output()
	relay1.High()
	relay2.Low()

	fmt.Println("Ready")
	for {
		count = 0

		if pin.Read() == 0 {
			if power == 0 {
				lastPinRead = 0
				pressCount = 1

				for {
					pinState = int(pin.Read())
					if pinState != lastPinRead {
						time.Sleep(200 * time.Millisecond) //debouncing
						pinState = int(pin.Read())

						lastPinRead = pinState

						if pinState == 0 {
							pressCount++
							count = 0
						}
					}

					time.Sleep(1 * time.Millisecond)
					count++
					if count == 600 {
						break
					}

				}

				if pressCount == 1 {

					fmt.Println("SINGLE")

					power = 1
					relay1.Low()
					relay2.High()
					cmdMusic := exec.Command("cvlc", "--random", "--play-and-exit", "/home/pi/Desktop/DiscoToilet/music/a_one_tap.mp3")
					cmdSpeaker := exec.Command("amixer", "-c", "2", "set", "Mic", "0%")
					fmt.Println("On")
					cmdMusicErr := cmdMusic.Start()
					cmdSpeakerErr := cmdSpeaker.Run()
					// time.Sleep(1000 * time.Millisecond)
					resetLights()
					turnOnLights()
					if cmdMusicErr != nil {
						fmt.Println(fmt.Sprint("Error: ", cmdMusicErr.Error()))
					}
					if cmdSpeakerErr != nil {
						fmt.Println(fmt.Sprint("Error: ", cmdSpeakerErr.Error()))
					}
					time.Sleep(500 * time.Millisecond)

					// }else if pressCount == 2{
					// 	cmdMusic := exec.Command("cvlc", "--random", "--play-and-exit", "/home/pi/Music/bathroom1.xspf")
					// 	_ = cmdMusic
					// 	fmt.Println("DOUBLE")

					// }else if pressCount < 2{
					// 	fmt.Println("TRIPPLE")

				} else {
					fmt.Println("DOUBLE OR MORE")

					power = 1
					relay1.Low()
					relay2.High()

					// time.Sleep(1000 * time.Millisecond)

					cmdMusic := exec.Command("cvlc", "--random", "--play-and-exit", "/home/pi/Desktop/DiscoToilet/music/a_two_tap.mp3")
					cmdSpeaker := exec.Command("amixer", "-c", "2", "set", "Mic", "0%")
					fmt.Println("On")
					cmdMusicErr := cmdMusic.Start()
					cmdSpeakerErr := cmdSpeaker.Run()

					resetLights()
					turnOnLightsBossMode()

					if cmdMusicErr != nil {
						fmt.Println(fmt.Sprint("Error: ", cmdMusicErr.Error()))
					}
					if cmdSpeakerErr != nil {
						fmt.Println(fmt.Sprint("Error: ", cmdSpeakerErr.Error()))
					}
					time.Sleep(500 * time.Millisecond)

				}
			}

			pinState = int(pin.Read())

			if power == 1 && pinState == 0 {
				power = 0
				relay1.High()
				relay2.Low()
				cmdMusic := exec.Command("killall", "vlc")
				cmdSpeaker := exec.Command("amixer", "-c", "2", "set", "Mic", "100%")
				fmt.Println("Off")

				resetLights()
				cmdMusicErr := cmdMusic.Run()
				cmdSpeakerErr := cmdSpeaker.Run()

				if cmdMusicErr != nil {
					fmt.Println(fmt.Sprint("Error: ", cmdMusicErr.Error()))
				}
				if cmdSpeakerErr != nil {
					fmt.Println(fmt.Sprint("Error: ", cmdSpeakerErr.Error()))
				}
				time.Sleep(500 * time.Millisecond)
			}
		}
	}
}

/* ---- Old Code ------
time.Sleep(1400 * time.Millisecond)

cmdLights := exec.Command("go", "run", "/home/pi/Desktop/button/turn-on-lights.go")
goErr := cmdLights.Run()
if goErr != nil {
	fmt.Println(fmt.Sprint("Error: ", goErr.Error()))
}

cmdLights := exec.Command("go", "run", "/home/pi/Desktop/button/reset.go")
goErr := cmdLights.Run()
if goErr != nil {
	fmt.Println(fmt.Sprint("Error: ", err.Error()))
}

//Sphere
	dmx.SetChannel(1, 255) //dimmer
	dmx.SetChannel(2, 0)   //strobe
	dmx.SetChannel(3, 255) //red
	dmx.SetChannel(4, 0)   //green
	dmx.SetChannel(5, 255) //blue
	dmx.SetChannel(6, 255) //rotate

	//Cage
	dmx.SetChannel(7, 255)  //dimmer
	dmx.SetChannel(8, 0)    //strobe
	dmx.SetChannel(9, 255)  //red
	dmx.SetChannel(10, 0)   //green
	dmx.SetChannel(11, 255) //blue
	dmx.SetChannel(12, 255) //white
	dmx.SetChannel(13, 255) //rotate

	//Middle Light
	dmx.SetChannel(14, 100) //dimmer
	dmx.SetChannel(15, 0)   //strobe
	dmx.SetChannel(16, 0)   //red
	dmx.SetChannel(17, 30)  //green
	dmx.SetChannel(18, 0)   //blue
	dmx.SetChannel(19, 0)   //yellow

	//Laser
	dmx.SetChannel(20, 125) //laser r
	dmx.SetChannel(21, 0)   //laser g
	dmx.SetChannel(22, 140) //animation & rotaion speed
	//dmx.SetChannel(23, 255)//???????

	//24.11.20
	//Sphere
	dmx.SetChannel(1, 255) //dimmer
	dmx.SetChannel(2, 0)   //strobe
	dmx.SetChannel(3, 255) //red
	dmx.SetChannel(4, 0)   //green
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
	dmx.SetChannel(14, 100) //dimmer
	dmx.SetChannel(15, 0)   //strobe
	dmx.SetChannel(16, 0)   //red
	dmx.SetChannel(17, 30)  //green
	dmx.SetChannel(18, 0)   //blue
	dmx.SetChannel(19, 0)   //yellow

	//Laser
	dmx.SetChannel(20, 0) //laser r 125
	dmx.SetChannel(21, 0) //laser g
	dmx.SetChannel(22, 0) //animation & rotaion speed 140
	// dmx.SetChannel(23, 180) //???????

	//Light #2
	dmx.SetChannel(24, 40)  // Auto Program
	dmx.SetChannel(25, 180) // Program Speed
	dmx.SetChannel(26, 255) // Manual Color Option
	dmx.SetChannel(27, 255) // Strobe
	dmx.SetChannel(28, 255) // Motor
	dmx.SetChannel(29, 255) // Pattern Options
	dmx.SetChannel(30, 0)   //Laser Colors
	dmx.SetChannel(31, 0)   //Laser Strobes
	dmx.SetChannel(32, 0)   //Laser Rotation

	dmx.SetChannel(33, 255)
	dmx.SetChannel(34, 255)
	dmx.SetChannel(35, 255)
	dmx.SetChannel(36, 255)
	dmx.SetChannel(37, 255)
	dmx.SetChannel(38, 255)


	//Sphere
	dmx.SetChannel(1, 255) //dimmer
	dmx.SetChannel(2, 0)   //strobe
	dmx.SetChannel(3, 255) //red
	dmx.SetChannel(4, 0)   //green
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
	dmx.SetChannel(14, 100) //dimmer
	dmx.SetChannel(15, 0)   //strobe
	dmx.SetChannel(16, 0)   //red
	dmx.SetChannel(17, 30)  //green
	dmx.SetChannel(18, 0)   //blue
	dmx.SetChannel(19, 0)   //yellow

	//Laser
	dmx.SetChannel(20, 0) //laser r 125
	dmx.SetChannel(21, 0) //laser g
	dmx.SetChannel(22, 0) //animation & rotaion speed 140
	// dmx.SetChannel(23, 180) //???????

	dmx.SetChannel(1, 255) //dimmer
	dmx.SetChannel(2, 0)   //strobe
	dmx.SetChannel(3, 255) //red
	dmx.SetChannel(4, 0)   //green
	dmx.SetChannel(5, 255) //blue
	dmx.SetChannel(6, 80)  //moon rotate
	dmx.SetChannel(7, 0)   //dimmer
	dmx.SetChannel(8, 0)   //strobe
	dmx.SetChannel(9, 0)   //red
	dmx.SetChannel(10, 0)  //green
	dmx.SetChannel(11, 0)  //blue
	dmx.SetChannel(12, 0)  //yellow
	dmx.SetChannel(13, 0)  //laser r
	dmx.SetChannel(14, 0)  //laser g
	dmx.SetChannel(15, 0)  //pan rotate

	-- End --

 --------------------- */
