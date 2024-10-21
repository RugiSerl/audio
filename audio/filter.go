package audio

import (
	"audio/math"
	"audio/utils"
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

// fourrier filter
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

// RC filter
func LowPassFilterTest(data *audio.IntBuffer, smoothness float64) *audio.IntBuffer {
	var position float64 = 0

	for i := range data.Data {
		position = float64(data.Data[i])*(1-smoothness) - position*smoothness
		data.Data[i] = int(position)
	}

	return data
}

// just a hard clipper
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

// pretty straightforward
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

// Not working, naive way
func Compressor(data *audio.IntBuffer) *audio.IntBuffer {
	data = Normalize(data)
	amplitudeMax := float64(math.PowInt(2, data.SourceBitDepth-1) - 1)
	var Weekness float64 = 100 // need > 0. Makes the function converge to Id when weekness -> infinity

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

// works as intended
func Delay(data *audio.IntBuffer) *audio.IntBuffer {
	const (
		INTERVAL  float64 = 0.0005 //s
		FADE_TIME float64 = 0.005  //s
		VOLUME    float64 = 1e-1
	)

	var intervalSamples float64 = float64(data.Format.SampleRate) * INTERVAL
	var fadeTimeSamples float64 = float64(data.Format.SampleRate) * FADE_TIME

	fmt.Println(fadeTimeSamples / intervalSamples)

	data_float64 := utils.Map(data.Data, func(i int) float64 {
		return float64(i)
	})

	for i := range data.Data {
		for j := 0; j < int(fadeTimeSamples); j += int(intervalSamples) {
			if i+j < len(data_float64) {
				data_float64[i+j] += float64(data.Data[i]) * (fadeTimeSamples - float64(j)) * VOLUME
			}
		}
	}
	data.Data = utils.Map(data_float64, func(f float64) int {
		return int(f)
	})
	data = Normalize(data)

	return data
}

// Yes it works, what is the problem ?
func Stretch(data *audio.IntBuffer, factor float64) *audio.IntBuffer {

	data.Format.SampleRate = int(float64(data.Format.SampleRate) * factor)

	return data
}

// Pretty ambitious
// How it should work : https://lh4.googleusercontent.com/bPyO-3yO0eMMfyQ_vE0SlM2UYrILdjCqipl5tePp4aNUmnzrwU03i-fQ2PjWTOH3aW5mPNsdVdxA_Wt11ur_6vfqQE-b448EmQXjJ36MC16hSCyieIu7A7p0ukzF_FYQCA=w1280
func Paulstretch(data *audio.IntBuffer, factor float64) *audio.IntBuffer {
	const (
		WINDOW_SIZE float64 = 0.1  //s
		STEP        float64 = 0.03 //s. Be careful, STEP/factor < WINDOW_SIZE
	)
	var windowSizeSamples float64 = float64(data.Format.SampleRate) * WINDOW_SIZE
	var stepSamples float64 = float64(data.Format.SampleRate) * STEP
	var data_float64 = utils.Map(data.Data,
		func(i int) float64 {
			return float64(i)
		},
	)
	var data_dest = make([]float64, int(float64(len(data_float64))/factor))

	var windowFunc = func(x float64) float64 {
		if x < -1 && x > 1 {
			return 0
		} else {
			return m.Pow(1-x*x, 1.25)
		}
	}
	fmt.Println("data_float64", len(data_float64))
	fmt.Println("data_dest", len(data_dest))
	for step := 0; step*int(stepSamples) < len(data.Data)-int(windowSizeSamples)-1; step++ {
		for i := 0; i < int(windowSizeSamples)-1; i++ {
			data_dest[int(float64(step)*stepSamples/factor)+i] += data_float64[int(float64(step)*stepSamples)+i] * windowFunc(float64(i)/windowSizeSamples*2-1)
		}
	}

	data.Data = utils.Map(data_dest,
		func(f float64) int {
			return int(f)
		},
	)
	data = Normalize(data)

	return data
}
