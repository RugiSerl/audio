package main

import "audio/audio"

func main() {
	buf, data, _ := audio.Parse("assets/cloches.wav")
	audio.FourierTest(buf.Data, uint16(buf.SourceBitDepth))

	audio.Save("assets/output.wav", buf, data)

}
