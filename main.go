package main

import (
	"audio/audio"
	"audio/math"
	"audio/utils"
	"audio/window"
	"audio/window/visual"
	"fmt"
)

const (
	AUDIO_FILENAME = "assets/cloches.wav"
	// AUDIO_FILENAME = "assets/amen_break/cw_amen02_165.wav"
)

func demo() {
	fmt.Println("Parsing audio...")
	buf, data, _ := audio.Parse(AUDIO_FILENAME)
	fmt.Println("Parsing done.")
	fmt.Println(buf.Format.NumChannels, buf.Format.SampleRate)
	fmt.Println("this is the parameters : ", data)
	visual.GenerateImage(math.WaveletTransform(math.HareWave, utils.Map(buf.Data, func(e int) float64 { return float64(e) })), "wavelet")

	buf = audio.Filter(buf, 600)
	fmt.Println("Filtering done. Saving...")
	audio.Save("assets/output.wav", buf, data)

	fmt.Println("Saving done.")
	list := visual.FourierTest(buf)
	fmt.Println("spectre done, initializing window")

	window.DisplayVisual(list, "assets/output.wav")
	//window.Test()

}

func main() {
	format := audio.GetDefaultFormat()
	bitDepth := 24
	bufGenerated := audio.GenerateWavetable(format, 5, 110, bitDepth, audio.SineWavetable)
	audio.Save("assets/generated.wav", bufGenerated, audio.Parameters{
		WaveFormatType: 1,
		Channel:        1,
		SampleRate:     format.SampleRate,
		BitsPerSample:  bitDepth,
	})
	fmt.Println("Done.")
}
