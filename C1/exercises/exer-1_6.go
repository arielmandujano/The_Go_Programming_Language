package exercises

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

// Modify de Lissajous program to produce images in multiple colors by adding
// more values to palette and then displaying them by changing the third argument
// of Set-ColorIndex in some intersting way

var palette2 = []color.Color{color.White, color.RGBA{0xb0, 0x95, 0xe2, 0xFF}, color.RGBA{0xd6, 0x95, 0xe2, 0xFF}, color.RGBA{0xe2, 0x95, 0xc7, 0xFF}, color.RGBA{0xe2, 0x95, 0xa1, 0xFF}, color.RGBA{0xe2, 0xb0, 0x95, 0xFF}, color.RGBA{0xe2, 0xd6, 0x95, 0xFF}}

func Exer1_6(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animations frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase diference
	lastUsed := 0
	for i := 0; i < nframes; i++ {
		lastUsed = getColorIndex(lastUsed, len(palette2))
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette2)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(lastUsed))
		}
		phase += 1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func getColorIndex(lastUsed int, colorCounts int) int {
	if lastUsed == colorCounts {
		return 1
	}
	return lastUsed + 1
}
