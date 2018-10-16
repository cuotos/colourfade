package main

import (
	"fmt"
	"github.com/aybabtme/rgbterm"
	"time"
)

var (
	rgb = [3]uint8{255, 0, 0}
	increment uint8 = 1
	interval = time.Millisecond * 10
	limiter chan struct{}
)


func main() {
	limiter = make(chan struct{})

	go func() {
		for range time.NewTicker(interval).C {
			limiter <- struct{}{}
		}
	}()

	dec := 0

	for {
		inc := (dec + 1) % 3

		for rgb[inc] < 255 {
			rgb[inc] = rgb[inc] + increment
			<-limiter
			print(rgb)
		}

		for rgb[dec] > 0 {
			rgb[dec] = rgb[dec] - increment
			<-limiter
			print(rgb)
		}

		dec = (dec + 1) % 3
	}
}

func print(rgb [3]uint8) {
	s := fmt.Sprintf("R: %03d, G: %03d, B: %03d - #%02[1]x%02[2]x%02[3]x\n", rgb[0], rgb[1], rgb[2])

	coloredWord := rgbterm.FgString(s, rgb[0], rgb[1], rgb[2])
	fmt.Print(coloredWord)
}
