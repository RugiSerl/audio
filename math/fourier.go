package math

import "math"

//list of magnitudes of complex numbers, with length = n of frequency bins
type Magnitudes []float64

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

func Ftransform(sample []int, freqbin int) Complex {
	N := len(sample)
	sum := Complex{Re: 0, Im: 0}
	for n := 0; n < N; n++ {
		sum = Add(sum, Mult(Complex{Re: float64(sample[n]), Im: 0}, ExpI(-2*math.Pi*float64(freqbin)*float64(n)/float64(N))))
	}
	return sum
}
