package main

import (
	"audio/audio"
	"audio/window"
)

const (
	AUDIO_FILENAME = "assets/cloches.wav"
)

func main() {
	buf, data, _ := audio.Parse(AUDIO_FILENAME)
	list := audio.FourierTest(buf.Data, uint16(buf.SourceBitDepth))
	window.InitVisual(buf, list)
	audio.Save("assets/output.wav", buf, data)

}
