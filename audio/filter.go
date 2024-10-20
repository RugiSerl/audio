package audio

import (
	"audio/math"
	"fmt"
	m "math"

	"github.com/go-audio/audio"
)

// Analog filter to image blur
func Blur(data *audio.IntBuffer) *audio.IntBuffer {
	newData := initBufferWithOthersStat(data)
	var kerSize int = 4

	for i := 0; i < len(data.Data); i++ {
		avg := 0
		for j := i - kerSize; j <= i+kerSize; j++ {
			avg += int(float32(data.Data[i]) / (2*float32(kerSize) + 1))
		}

		newData.Data[i] = avg
	}
	return newData

}

// very simple filter changing the amplitude of the audio
func ReduceAmp(data *audio.IntBuffer) *audio.IntBuffer {
	var factor float32 = 0.1
	for i := 0; i < len(data.Data); i++ {
		data.Data[i] = int(float32(data.Data[i]) * factor)
	}
	return data
}

func PassFilter(fourierCoefficients math.FrequencyDomainData, freqOffset float64, ease float64) math.FrequencyDomainData {
	// frequency domain manipulation
	for i := 0; i < len(fourierCoefficients); i++ {
		fourierCoefficients[i] = math.Mult(fourierCoefficients[i], math.Real(1/(1+m.Exp((1/ease)*(float64(float64(i)-float64(len(fourierCoefficients))*freqOffset))))))
	}

	return fourierCoefficients
}

func Filter(data *audio.IntBuffer, freqLimit int) *audio.IntBuffer {
	TimeInterval := int(m.Pow(2, m.Ceil(m.Log2(float64(len(data.Data)))))) //samples

	// Signal decomposition
	fourrierCoefficients := math.FFT(math.MapIntArrayToTimeDomainData(math.AddZeroPadding(data.Data, TimeInterval)))

	// Frequency filtering
	fourrierCoefficients = PassFilter(fourrierCoefficients, 0.995, -0.1)

	// Reconstruction of signal
	data.Data = math.MapTimeDomainDataToIntArray(math.InverseFFT(fourrierCoefficients))

	return data
}

func LowPassFilterTest(data *audio.IntBuffer, smoothness float64) *audio.IntBuffer {
	var position float64 = 0

	for i := range data.Data {
		position = float64(data.Data[i])*(1-smoothness) - position*smoothness
		data.Data[i] = int(position)
	}

	return data
}

func Limiter(data *audio.IntBuffer) *audio.IntBuffer {
	const gain float64 = 2
	fmt.Println(data.SourceBitDepth)
	amplitudeMax := float64(math.PowInt(2, data.SourceBitDepth-1) - 1)
	var value float64

	for i := range data.Data {
		value = gain * float64(data.Data[i])

		if m.Abs(value) > amplitudeMax {
			value *= amplitudeMax / m.Abs(value)
		}

		data.Data[i] = int(value)

	}

	return data

}

func Normalize(data *audio.IntBuffer) *audio.IntBuffer {
	amplitudeMax := float64(math.PowInt(2, data.SourceBitDepth-1) - 1)
	var Max float64 = 0
	for i := range data.Data {
		Max = max(Max, m.Abs(float64(data.Data[i])))
	}
	for i := range data.Data {
		data.Data[i] = int(amplitudeMax * float64(data.Data[i]) / Max)
	}

	return data

}

func Compressor(data *audio.IntBuffer) *audio.IntBuffer {
	data = Normalize(data)
	amplitudeMax := float64(math.PowInt(2, data.SourceBitDepth-1) - 1)
	var Weekness float64 = 200 // need > 0. Makes the function converge to Id when weekness -> infinity

	for i := range data.Data {
		x := float64(data.Data[i])
		if data.Data[i] > 0 {
			data.Data[i] = int((amplitudeMax * (m.Pow(x/amplitudeMax-1, 2) + 1 + Weekness*x/amplitudeMax) / (Weekness + 1)))
		} else {
			data.Data[i] = int((-amplitudeMax * (m.Pow(-x/amplitudeMax-1, 2) + 1 - Weekness*x/amplitudeMax) / (Weekness + 1)))

		}
	}

	return data
}
