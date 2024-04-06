package window

import (
	"audio/window/visual"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/go-audio/audio"
)

const (
	maxSamplesPerUpdate = 4096
	SPECTRE_WIDTH       = 400 //px
	SPECTRE_HEIGHT      = 200 //px

)

var (
	soundStart float32 = 0
)

func displayVisual(m visual.MagnitudesList, time float32) {
	i := int((time - soundStart) / m.DeltaTime)
	if i < len(m.Data) {
		rl.DrawRectangleV(rl.NewVector2(float32(rl.GetScreenWidth())/2-SPECTRE_WIDTH/2, float32(rl.GetScreenHeight())/2-SPECTRE_HEIGHT/2), rl.NewVector2(SPECTRE_WIDTH, SPECTRE_HEIGHT), rl.Gray)

		for j := 0; j < len(m.Data[i]); j++ {
			rl.DrawRectangleV(rl.NewVector2(float32(rl.GetScreenWidth())/2-SPECTRE_WIDTH/2+float32(j)*SPECTRE_WIDTH/float32(len(m.Data[i])+1), float32(rl.GetScreenHeight())/2-SPECTRE_HEIGHT/2), rl.NewVector2(SPECTRE_WIDTH/float32(len(m.Data)), SPECTRE_HEIGHT*float32(m.Data[i][j])), rl.Black)

		}

	}

}

// temp
func playAudio(filename string) {
	m := rl.LoadSound(filename)

	rl.PlaySound(m)
	fmt.Println(m.FrameCount)
	soundStart = float32(rl.GetTime())
}

func InitVisual(audioBuffer *audio.IntBuffer, magnitudeList visual.MagnitudesList, filename string) {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	rl.InitAudioDevice()

	playAudio(filename)

	//load all the dat
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		displayVisual(magnitudeList, float32(rl.GetTime()))

		rl.EndDrawing()
	}

	rl.CloseAudioDevice() // Close audio device (music streaming is automatically stopped)

	rl.CloseWindow()
}
