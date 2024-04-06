package main

import (
	"audio/audio"
	"audio/window"
	"audio/window/visual"
	"fmt"
)

const (
	AUDIO_FILENAME = "assets/cloches.wav"
	// AUDIO_FILENAME = "assets/amen_break/cw_amen02_165.wav"
)

func main() {
	fmt.Println("parsing audio...")
	buf, data, _ := audio.Parse(AUDIO_FILENAME)
	fmt.Println("parsing done.")

	buf = audio.Filter(buf, 600)
	audio.Save("assets/output.wav", buf, data)
	list := visual.FourierTest(buf)

	window.InitVisual(buf, list, "assets/output.wav")
	//window.Test()

}
