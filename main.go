package main

import (
	"fmt"
	"github.com/aybabtme/rgbterm"
	"time"
)

var (
	rgb = [3]uint8{255, 0, 0}
	interval = time.Millisecond * 20
)

func main() {

	dec := 0
	inc := 1

	for {
		for rgb[inc] < 255 {
			rgb[inc] = rgb[inc] + 5
			print(rgb)
		}

		for rgb[dec] > 0 {
			rgb[dec] = rgb[dec] - 5
			print(rgb)
		}

		inc = (inc + 1) % 3
		dec = (dec +1) % 3
	}

}

func print(rgb [3]uint8) {
	s := fmt.Sprintf("R: %03d, G: %03d, B: %03d - #%02[1]x%02[2]x%02[3]x\n", rgb[0], rgb[1], rgb[2])

	coloredWord := rgbterm.FgString(s, rgb[0], rgb[1], rgb[2])
	fmt.Print(coloredWord)
	time.Sleep(interval)
}