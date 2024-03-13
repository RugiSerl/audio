package plot

import (
	"audio/math"
	"image"
	"image/color"
	"image/png"
	"os"
)

func GenerateImage(data []math.Magnitudes, name string) {

	width := len(data)
	height := len(data[0])

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			intensity := uint8(data[x][y] * 255)
			img.Set(x, height-y, color.RGBA{0, intensity, intensity, 0xff})
		}
	}

	// Encode as PNG.
	f, _ := os.Create("assets/" + name + ".png")
	png.Encode(f, img)

}
