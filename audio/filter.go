package audio

import (
	"audio/math"
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
	fourrierCoefficients = PassFilter(fourrierCoefficients, .9, -.01)

	// Reconstruction of signal
	data.Data = math.MapTimeDomainDataToIntArray(math.InverseFFT(fourrierCoefficients))

	return data
}

func LowPassFilterTest(data *audio.IntBuffer, freqLimit int) *audio.IntBuffer {
	const tau = 60
	var position float64 = 0

	for i := range data.Data {
		position = (float64(data.Data[i]) - position) / tau
		data.Data[i] = int(position)
	}

	return data
}
