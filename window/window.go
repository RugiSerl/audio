package window

import (
	m "audio/math"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/go-audio/audio"
)

const (
	maxSamplesPerUpdate = 4096
)

func InitVisual(audioBuffer *audio.IntBuffer, magnitudeList []m.Magnitudes) {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	rl.InitAudioDevice()
	stream := rl.LoadAudioStream(uint32(audioBuffer.Format.SampleRate), uint32(audioBuffer.SourceBitDepth), uint32(audioBuffer.Format.NumChannels))

	//load all the data
	data := make([]float32, len(audioBuffer.Data))
	maxAmplitude := float32(math.Pow(2, float64(audioBuffer.SourceBitDepth)))
	for i := 0; i < len(audioBuffer.Data); i++ {
		data[i] = float32(audioBuffer.Data[i]) / maxAmplitude
	}

	// NOTE: The generated MAX_SAMPLES do not fit to close a perfect loop
	// for that reason, there is a clip everytime audio stream is looped
	rl.PlayAudioStream(stream)

	totalSamples := int32(len(audioBuffer.Data))
	samplesLeft := int32(totalSamples)

	rl.SetTargetFPS(30)

	for !rl.WindowShouldClose() {
		// Refill audio stream if required
		if rl.IsAudioStreamProcessed(stream) {
			numSamples := int32(0)
			if samplesLeft >= maxSamplesPerUpdate {
				numSamples = maxSamplesPerUpdate
			} else {
				numSamples = samplesLeft
			}

			rl.UpdateAudioStream(stream, data[totalSamples-samplesLeft:])

			samplesLeft -= numSamples

			// Reset samples feeding (loop audio)
			if samplesLeft <= 0 {
				samplesLeft = totalSamples
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.EndDrawing()
	}

	rl.UnloadAudioStream(stream) // Close raw audio stream and delete buffers from RAM

	rl.CloseAudioDevice() // Close audio device (music streaming is automatically stopped)

	rl.CloseWindow()
}
