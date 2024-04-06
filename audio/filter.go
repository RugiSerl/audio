package audio

import (
	"audio/math"
	"fmt"

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

func Filter(data *audio.IntBuffer, freqLimit int) *audio.IntBuffer {
	TimeInterval := 65536 //samples
	var fourrierCoefficients []math.Complex

	for i := 0; i < len(data.Data)/TimeInterval; i++ {
		fmt.Println(i*TimeInterval, " samples")

		//signal decomposition
		fourrierCoefficients = math.FFT(math.MapIntArrayToTimeDomainData(data.Data[i*TimeInterval : (i+1)*TimeInterval]))

		// frequency domain manipulation
		// for i := len(fourrierCoefficients) / 8; i < len(fourrierCoefficients); i++ {
		// 	fourrierCoefficients[i] = math.Real(0)
		// }

		//reconstruction of signal
		sample := math.MapTimeDomainDataToIntArray(math.InverseFFT(fourrierCoefficients))
		for j := 0; j < TimeInterval; j++ {
			data.Data[i*TimeInterval+j] = sample[j]
		}

	}

	return data
}
