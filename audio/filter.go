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

func FourierTest(data []int, bitDepth uint16) []math.Magnitudes {
	TimeInterval := 500 //samples
	List := []math.Magnitudes{}
	var maxMagnitude float64 = 0 // positive real number

	for i := 0; i < len(data)/TimeInterval; i++ {
		fmt.Println(i*TimeInterval, " samples")
		magnitudes, localmax := math.GetMagnitudes(data, i*TimeInterval, (i+1)*TimeInterval)
		List = append(List, magnitudes)

		if localmax > maxMagnitude {
			maxMagnitude = localmax
		}

	}

	for i := range List {
		for j := range List[i] {
			//we are normalizing all the magnitudes to get everything between 0 and 1
			List[i][j] /= maxMagnitude

		}
	}

	plot.GenerateImage(List, "fourrier")

	return List
}
