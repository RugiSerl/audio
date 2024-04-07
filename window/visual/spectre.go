package visual

import (
	"audio/math"

	"github.com/go-audio/audio"
)

// list of magnitudes of complex numbers, with length = n of frequency bins
type Magnitudes []float64
type MagnitudesList struct {
	Data         []Magnitudes
	SampleAmount int
	DeltaTime    float32
}

func GetMagnitudes(data math.TimeDomainData, begin int, end int) (Magnitudes, float64) {
	var maxMagnitude float64 = 0 // positive real number
	magnitudes := Magnitudes{}
	coefs := math.FFT(data[begin:end])
	for k := 1; k < (end-begin)/2; k++ {
		mag := math.Norm(coefs[k])
		magnitudes = append(magnitudes, mag)
		if mag > maxMagnitude {
			maxMagnitude = mag
		}
	}
	return magnitudes, maxMagnitude
}

func FourierTest(data *audio.IntBuffer) MagnitudesList {
	TimeInterval := 512 //samples
	List := MagnitudesList{}
	var maxMagnitude float64 = 0 // positive real number

	for i := 0; i < len(data.Data)/TimeInterval; i++ {
		// fmt.Println(i*TimeInterval, " samples")
		magnitudes, localmax := GetMagnitudes(math.MapIntArrayToTimeDomainData(data.Data), i*TimeInterval, (i+1)*TimeInterval)
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

	GenerateImage(List.Data, "fourrier")

	List.SampleAmount = len(data.Data) / TimeInterval
	List.DeltaTime = 1 / (float32(data.Format.SampleRate) / float32(TimeInterval))

	return List
}
