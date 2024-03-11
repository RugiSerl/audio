package audio

import (
	"fmt"
	"math"
	"os"

	"github.com/zenwerk/go-wave"
)

type Parameters struct {
	WaveFormatType int
	Channel        int
	SampleRate     int
	BitsPerSample  int
}

// raw data containing the samples data
type PcmBuffer [][]float64

// parse all the data all at once. Simple but unoptimised
func Parse(fileName string) PcmBuffer {
	reader, err := wave.NewReader(fileName)

	if err != nil {
		panic(err)
	}

	var e error = nil
	var s PcmBuffer = make(PcmBuffer, reader.NumSamples)

	for i := 0; e == nil; i++ {
		s[i], e = reader.ReadSample()
		fmt.Println(s[i])
	}

	return s

}

// Save the whole buffer
func Save(fileName string, buffer PcmBuffer, param Parameters) {
	f, err := os.Create(fileName)

	parameter := wave.WriterParam{
		Out:            f,
		WaveFormatType: param.WaveFormatType,
		Channel:        param.Channel,
		SampleRate:     param.SampleRate,
		BitsPerSample:  param.BitsPerSample,
	}

	if err != nil {
		panic(err)
	}

	w, err := wave.NewWriter(parameter)

	for _, e := range buffer {
		w.WriteSample16([]int16{int16(e[0] * math.Pow(2, 24))})

	}

	if err != nil {
		panic(err)
	}

	w.Close()

}
