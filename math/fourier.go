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

type timeDomainData = []float64
type FrequencyDomainData = []Complex

// Annoying practical function because most of audio format samples are of type int
func MapIntArrayToTimeDomainData(samples []int) timeDomainData {
	r := make(timeDomainData, len(samples))
	for i := 0; i < len(samples); i++ {
		r[i] = float64(samples[i])
	}
	return r
}

// Annoying practical function because most of audio format samples are of type int
func MapTimeDomainDataArrayToInt(samples timeDomainData) []int {
	r := make([]int, len(samples))
	for i := 0; i < len(samples); i++ {
		r[i] = int(samples[i])
	}
	return r
}

func GetMagnitudes(data timeDomainData, begin int, end int) (Magnitudes, float64) {
	var maxMagnitude float64 = 0 // positive real number
	magnitudes := Magnitudes{}
	coefs := Ftransform(data[begin:end])
	for k := 1; k < (end-begin)/2; k++ {
		mag := Norm(coefs[k])
		magnitudes = append(magnitudes, mag)
		if mag > maxMagnitude {
			maxMagnitude = mag
		}
	}
	return magnitudes, maxMagnitude
}

// Naive DTF algorithm in O(n²) (nested loops).
// Take time domain samples and returns frequency domains values.
// The values in the frequency domain are in that form :
// {amount of cosine of frequency} + i{amount of sine of frequency}.
// So to get the magnitude of the signal of the frequency a + ib,
// We calculate √(a²+b²) (Acos(wt) + Bsin(wt) = |A+iB|cos(wt + phi)).
// So phi ≡ artan(b/a) [pi]
func Ftransform(samples timeDomainData) FrequencyDomainData {
	N := len(samples)
	fourrierCoefficients := make([]Complex, N)
	//loop trough all frequencies possibilities
	for f := 0; f < N; f++ {

		//c[f] = ∑s[n]exp(-2πnf/N), n∈[0, N[
		for n := 0; n < N; n++ {
			fourrierCoefficients[f] = Add(fourrierCoefficients[f], Mult(Real(samples[n]), Omega(float64(f)*float64(n)/float64(N))))
		}
	}

	return fourrierCoefficients
}

// Inverse DFT, O(n²) and takes frequency domain coefficients to returns its time domains samples
// Very similar to DFT, since it is the same as calculating the inverse Vandermonde matrix
// (more information here : https://fr.wikipedia.org/wiki/Transformation_de_Fourier_rapide#Formulation_math%C3%A9matique)
func InverseFtransform(coefficients FrequencyDomainData) timeDomainData {
	N := len(coefficients)
	Samples := make([]float64, N)

	//loop trough all sample positions
	for n := 0; n < N; n++ {
		sum := Complex{0, 0}
		//s[n]*N = ∑c[f]exp(-2πnf/N), f∈[0, N[
		for f := 0; f < N; f++ {
			sum = Add(sum, Mult(coefficients[f], Omega(-1*float64(n)*float64(f)/float64(N))))
		}

		//reconstruct signal using the coefficients of sine and cosine
		Samples[n] = (1.0 / float64(N)) * (sum.Re*math.Cos(2*math.Pi*float64(n)) + sum.Im*math.Sin(2*math.Pi*float64(n)))

	}

	return Samples
}

// Cooley Tuckey divide and conquer algorithm.
// Information and pseudo code here : https://fr.wikipedia.org/wiki/Transformation_de_Fourier_rapide#Pseudo-code
func FFTAux(samples Polynomial, unityRoot Complex) FrequencyDomainData {
	n := len(samples.coefs)
	//constant polynomial case
	if n == 1 {
		return []Complex{samples.Evaluate(Real(1))}

	} else {

		evenPart := samples.Even()
		oddPart := samples.Odd()

		evenResults := FFTAux(evenPart, unityRoot.Pow(2))
		oddResults := FFTAux(oddPart, unityRoot.Pow(2))
		results := make([]Complex, n)

		// fmt.Println("even and odds", samples, evenPart, oddPart)
		// fmt.Println(N, evenResults, oddResults)

		for k := 0; k < n/2; k++ {

			results[k] = Add(evenResults[k], Mult(unityRoot, oddResults[k]))

		}

		return results
	}

}

func FFT(samples timeDomainData) FrequencyDomainData {
	return FFTAux(NewPolynomialFromReal(samples), Omega(1/float64(len(samples))))
}

// func inverseFFT(coefficients FrequencyDomainData) timeDomainData {
// 	return FFTAux(Polynomial{coefficients}, Omega(-1/float64(len(coefficients))))
// }
