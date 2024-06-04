package math

import "math"

type Wave = func(float64, float64) float64
type WaveletResult = [][]float64

func scalarProduct(w Wave, data []float64, dilatation float64) float64 {
	s := 0.
	var t float64
	for i, element := range data {
		t = float64(i)/float64(len(data)) - .5
		s += element * w(t, dilatation)
	}
	return s
}

func MexicanHatWave(t, dilatation float64) float64 {
	// retranscription of https://en.wikipedia.org/wiki/Ricker_wavelet
	return 2 / (math.Sqrt(3*dilatation) * math.Pow(math.Pi, 1./4.)) * (1 - math.Pow(t/dilatation, 2)) * math.Exp(-1./2.*math.Pow(t/dilatation, 2))
}

func HareWave(t, dilatation float64) float64 {
	t /= dilatation
	if t >= 0 && t < 1./2. {
		return 1
	} else if t >= 1./2. && t < 1 {
		return -1
	} else {
		return 0
	}
}

func Sinc(t, dilatation float64) float64 {
	t /= dilatation
	if t == 0 {
		return 1
	} else {
		return math.Sin(t) / t
	}
}

func WaveletTransform(w Wave, data []float64) WaveletResult {
	sampleAmount := len(data) / 100
	dilatationMaxSize := 400

	// Init the array
	result := make(WaveletResult, sampleAmount)
	for i := range result {
		result[i] = make([]float64, dilatationMaxSize)
	}

	maximum := 0.

	for offset := 1; offset < sampleAmount; offset++ {
		for dilatation := 1; dilatation < dilatationMaxSize; dilatation++ {
			//t := offset * sampleSize
			result[offset][dilatation] = scalarProduct(w, data[offset:offset+dilatation], float64(dilatation))
			maximum = max(result[offset][dilatation], maximum)
		}
	}
	// Normalize everything
	for offset := 0; offset < sampleAmount; offset++ {
		for dilatation := 1; dilatation < dilatationMaxSize; dilatation++ {
			result[offset][dilatation] /= maximum
		}
	}
	return result

}
