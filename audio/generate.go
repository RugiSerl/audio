package audio

import (
	m "audio/math"
	"math"

	"github.com/go-audio/audio"
)

type Wavetable func(float64) float64

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

// The period and the amplitude of the function given must be 1
func GenerateWavetable(format *audio.Format, length int, frequency int, bitDepth int, wavetable Wavetable) *audio.IntBuffer {
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
