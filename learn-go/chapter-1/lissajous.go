// Генерация анимированного GIF из случайных фигур Лиссажу

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5 // количество полных колебаний x
		res = 0.001 // угловое разрешение
		size = 100
		nframes = 64 // количество кадров анимации
		delay = 8 // задержка между кадрами (единица - 10мс)
	)
	rand.Seed(time.Now().UnixNano())
	freq := rand.Float64() * 3.0 // относительная частота колебаний у
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size + 1, 2*size + 1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size + int(x*size + 0.5), size + int(y*size + 0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}