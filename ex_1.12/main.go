package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		cycles, err := strconv.Atoi(r.Form.Get("cycles"))
		if err != nil {
			http.Error(w, "invalid param 'cycles': "+err.Error(), http.StatusBadRequest)
			return
		}
		res, err := strconv.ParseFloat(r.Form.Get("res"), 64)
		if err != nil {
			http.Error(w, "invalid param 'res': "+err.Error(), http.StatusBadRequest)
			return
		}
		lissajous(w, float64(cycles), res, 100, 64, 8)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(
	out io.Writer,
	cycles float64, // 20  // number of complete x oscillator revolutions
	res float64, // 0.001 // angular resolution
	size int, // 100   // image canvas covers [-size..+size]
	nframes int, // 64    // number of animation frames
	delay int, // 8     // delay between frames in 10ms units
) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
