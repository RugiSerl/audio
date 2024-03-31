package audio

import (
	"audio/math"
	"audio/plot"
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
	TimeInterval := 1024 //samples
	var fourrierCoefficients []math.Complex

	for i := 0; i < len(data.Data)/TimeInterval; i++ {
		fmt.Println(i*TimeInterval, " samples")

		//signal decomposition
		fourrierCoefficients = math.Ftransform(math.MapIntArrayToTimeDomainData(data.Data[i*TimeInterval : (i+1)*TimeInterval]))
		// frequency domain manipulation
		// for i := len(fourrierCoefficients) / 8; i < len(fourrierCoefficients); i++ {
		// 	fourrierCoefficients[i] = math.Complex{Re: 0, Im: 0}
		// }

		//reconstruction of signal
		sample := math.MapTimeDomainDataArrayToInt(math.InverseFtransform(fourrierCoefficients))
		for j := 0; j < TimeInterval; j++ {
			data.Data[i*TimeInterval+j] = sample[j]
		}

	}

	return data
}

func FourierTest(data *audio.IntBuffer) math.MagnitudesList {
	TimeInterval := 500 //samples
	List := math.MagnitudesList{}
	var maxMagnitude float64 = 0 // positive real number

	for i := 0; i < len(data.Data)/TimeInterval; i++ {
		fmt.Println(i*TimeInterval, " samples")
		magnitudes, localmax := math.GetMagnitudes(math.MapIntArrayToTimeDomainData(data.Data), i*TimeInterval, (i+1)*TimeInterval)
		List.Data = append(List.Data, magnitudes)

		if localmax > maxMagnitude {
			maxMagnitude = localmax
		}

	}

	for i := range List.Data {
		for j := range List.Data[i] {
			//we are normalizing all the magnitudes to get everything between 0 and 1
			List.Data[i][j] /= maxMagnitude

		}
	}

	plot.GenerateImage(List.Data, "fourrier")

	List.SampleAmount = len(data.Data) / TimeInterval
	List.DeltaTime = 1 / (float32(data.Format.SampleRate) / float32(TimeInterval))

	return List
}
