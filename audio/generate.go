package audio

import (
	m "audio/math"
	"math"

	"github.com/go-audio/audio"
)

// The period of the function given must be 1 and the amplitude 2, centered in 1
type Wavetable func(float64) float64

type Sound func(float64) float64

func GetDefaultFormat() *audio.Format {
	return audio.FormatMono48000
}

func createIntBuffer(data []int, format *audio.Format, bitDepth int) *audio.IntBuffer {
	return &audio.IntBuffer{
		Format:         format,
		Data:           data,
		SourceBitDepth: bitDepth,
	}
}

// Not really working
func GenerateSiren(format *audio.Format, length int, frequency int, bitDepth int, shape Wavetable) *audio.IntBuffer {
	data := make([]int, length*format.SampleRate)
	amplitudeMax := m.PowInt(2, bitDepth-1) - 1
	const (
		frequencyModulationAmplitude = float64(3)
		frequencyModulationSpeed     = float64(1)
	)
	var oscillatingFrequency float64

	for i := 0; i < length*format.SampleRate; i++ {
		oscillatingFrequency = float64(frequency) + SineWavetable(frequencyModulationSpeed*float64(i)/float64(format.SampleRate))*frequencyModulationAmplitude
		data[i] = int(float64(amplitudeMax) * shape(oscillatingFrequency*float64(i)/float64(format.SampleRate)))
	}

	return createIntBuffer(data, format, bitDepth)

}

func GeneratePeriodicWavetable(format *audio.Format, length int, frequency int, bitDepth int, wavetable Wavetable) *audio.IntBuffer {
	data := make([]int, length*format.SampleRate)
	amplitudeMax := m.PowInt(2, bitDepth-1) - 1

	for i := 0; i < length*format.SampleRate; i++ {
		data[i] = int(float64(amplitudeMax) * wavetable(float64(frequency)*float64(i)/float64(format.SampleRate)))
	}

	return createIntBuffer(data, format, bitDepth)
}

func SineWavetable(x float64) float64 {
	return math.Sin(x * 2 * math.Pi)
}

func SquareWavetable(x float64) float64 {
	return float64(int(2*x)%2)*2 - 1
}

func SawToothWavetable(x float64) float64 {
	return (x-math.Floor(x))*2 - 1
}

func TriangleWaveTable(x float64) float64 {
	return 2/math.Pi*math.Acos(math.Cos(2*math.Pi*x)) - 1
}

func SharpTriangleWaveTable(x float64) float64 {
	return math.Pow(TriangleWaveTable(x), 3)
}
