package main

import (
	"audio/audio"
	"audio/window"
	"audio/window/visual"
	"fmt"
)

const (
	AUDIO_FILENAME = "assets/ImperialMarch60.wav"
	// AUDIO_FILENAME = "assets/amen_break/cw_amen02_165.wav"
)

func main() {
	fmt.Println("Parsing audio...")
	buf, data, _ := audio.Parse(AUDIO_FILENAME)
	fmt.Println("Parsing done.")

	buf = audio.Filter(buf, 600)
	fmt.Println("Filtering done. Saving...")
	audio.Save("assets/output.wav", buf, data)
	fmt.Println("Saving done.")
	list := visual.FourierTest(buf)
	fmt.Println("spectre done, initializing window")

	window.InitVisual(buf, list, "assets/output.wav")
	//window.Test()

}
