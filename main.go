package main

import (
	"audio/audio"
	"audio/window"
)

const (
	AUDIO_FILENAME = "assets/cheval.wav"
	//AUDIO_FILENAME = "assets/amen_break/cw_amen02_165.wav"
)

func main() {
	buf, data, _ := audio.Parse(AUDIO_FILENAME)
	list := audio.FourierTest(buf, uint16(buf.SourceBitDepth))
	window.InitVisual(buf, list, AUDIO_FILENAME)
	//window.Test()
	audio.Save("assets/output.wav", buf, data)

}
