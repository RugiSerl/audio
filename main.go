package main

import (
	"audio/audio"
	"audio/window"
)

const (
	AUDIO_FILENAME = "assets/cloches.wav"
	// AUDIO_FILENAME = "assets/amen_break/cw_amen02_165.wav"
)

func main() {
	buf, data, _ := audio.Parse(AUDIO_FILENAME)

	buf = audio.Filter(buf, 600)
	audio.Save("assets/output.wav", buf, data)
	list := audio.FourierTest(buf)

	window.InitVisual(buf, list, "assets/output.wav")
	//window.Test()

}
