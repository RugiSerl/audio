package main

import (
	"audio/audio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

const (
	AUDIO_FILENAME = "assets/piano_sample.wav"
	// AUDIO_FILENAME = "assets/amen_break/cw_amen02_165.wav"
)

func playSound(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	fmt.Println("Now playing sound", filename)

	<-done
}

func demo() {
	fmt.Println("Parsing audio...")
	buf, data, _ := audio.Parse(AUDIO_FILENAME)
	fmt.Println("Parsing done.")

	//visual.GenerateImage(math.WaveletTransform(math.HareWave, utils.Map(buf.Data, func(e int) float64 { return float64(e) })), "wavelet")

	buf = audio.Paulstretch(buf, 0.5)
	fmt.Println("Filtering done. Saving...")
	audio.Save("assets/output.wav", buf, data)

	fmt.Println("Saving done.")
	//list := visual.FourierTest(buf)
	//fmt.Println("spectre done, initializing window")

	//window.DisplayVisual(list, "assets/output.wav")
	//window.Test()

	playSound("assets/output.wav")

}

func gererateSound() {
	format := audio.GetDefaultFormat()
	bitDepth := 24
	filename := "assets/generated.wav"
	bufGenerated := audio.GeneratePeriodicWavetable(format, 10, 440, bitDepth, audio.SineWavetable)
	audio.Save(filename, bufGenerated, audio.Parameters{
		WaveFormatType: 1,
		Channel:        1,
		SampleRate:     format.SampleRate,
		BitsPerSample:  bitDepth,
	})
	fmt.Println("Done.")

	fmt.Println("Now playing sound")
	playSound(filename)
}

func main() {
	demo()
}
