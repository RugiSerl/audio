package main

import "audio/audio"

func main() {
	buf := audio.Parse("assets/preamble.wav")
	buf = audio.ChangeAmplitude(buf, 0.1)

}
