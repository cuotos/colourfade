package main

import (
	"fmt"
	"github.com/aybabtme/rgbterm"
	"time"
)

var (
	rgb = [3]uint8{255, 0, 0}
	interval = time.Millisecond * 10
)

func main() {

	dec := 0

	for {
		inc := (dec + 1) % 3

		for rgb[inc] < 255 {
			rgb[inc] = rgb[inc] + 5
			print(rgb)
		}

		for rgb[dec] > 0 {
			rgb[dec] = rgb[dec] - 5
			print(rgb)
		}

		dec = (dec + 1) % 3
	}

}

func print(rgb [3]uint8) {
	s := fmt.Sprintf("R: %03d, G: %03d, B: %03d - #%02[1]x%02[2]x%02[3]x\n", rgb[0], rgb[1], rgb[2])

	coloredWord := rgbterm.FgString(s, rgb[0], rgb[1], rgb[2])
	fmt.Print(coloredWord)
	time.Sleep(interval)
}