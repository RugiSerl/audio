package math

import (
	"math"
)

// list of magnitudes of complex numbers, with length = n of frequency bins
type Magnitudes []float64
type MagnitudesList struct {
	Data         []Magnitudes
	SampleAmount int
	DeltaTime    float32
}

func GetMagnitudes(data []int, begin int, end int) (Magnitudes, float64) {
	var maxMagnitude float64 = 0 // positive real number
	magnitudes := Magnitudes{}
	for k := 1; k < (end-begin)/2; k++ {
		mag := Norm(Ftransform(data[begin:end], k))
		magnitudes = append(magnitudes, mag)
		if mag > maxMagnitude {
			maxMagnitude = mag
		}

	}

	return magnitudes, maxMagnitude

}

func Ftransform(samples []int, freqbin int) Complex {
	N := len(samples)
	sum := Complex{Re: 0, Im: 0}
	for n := 0; n < N; n++ {
		sum = Add(sum, Mult(Complex{Re: float64(samples[n]), Im: 0}, ExpI(-2*math.Pi*float64(freqbin)*float64(n)/float64(N))))
	}
	return sum
}

func InverseFtransform(coefficients []Complex, index int, bitDepth int) float64 {
	N := len(coefficients)
	sum := Complex{Re: 0, Im: 0}
	for n := 0; n < N; n++ {
		sum = Add(sum, Mult(coefficients[n], ExpI(2*math.Pi*float64(index)*float64(n)/float64(N))))
	}

	return (sum.Re*math.Cos(2*math.Pi*float64(index)) + sum.Im*math.Sin(2*math.Pi*float64(index))) / float64(N)
}
