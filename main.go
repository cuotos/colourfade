package main

import (
	"fmt"
	"github.com/aybabtme/rgbterm"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

var (
	rgb             = [3]uint8{255, 0, 0}
	increment uint8 = 1
	interval        = time.Millisecond * 10
	limiter         = make(<-chan time.Time)
)

func main(){
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("colour fade", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
}

func old() {
	limiter = time.NewTicker(interval).C

	dec := 0

	for {
		inc := (dec + 1) % 3
		colourCycle(dec, inc)
		dec = (dec + 1) % 3
	}
}

func colourCycle(dec, inc int) {

	for rgb[inc] < 255 {
		rgb[inc] += increment
		<-limiter
		print(rgb)
	}

	for rgb[dec] > 0 {
		rgb[dec] -= increment
		<-limiter
		print(rgb)
	}
}

func print(rgb [3]uint8) {
	s := fmt.Sprintf("R: %03d, G: %03d, B: %03d - #%02[1]x%02[2]x%02[3]x\n", rgb[0], rgb[1], rgb[2])

	coloredWord := rgbterm.FgString(s, rgb[0], rgb[1], rgb[2])

	fmt.Print(coloredWord)
}
