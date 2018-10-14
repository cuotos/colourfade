package main

import (
	"fmt"
	"github.com/aybabtme/rgbterm"
	"math/rand"
	"time"
)

var (
	rgb = [3]uint8{0, 0, 0}
	interval = time.Millisecond * 20
	r *rand.Rand
)

func init(){
	r = rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range rgb {
		if r.Intn(2) == 1 {
			rgb[i] = 255
		}
	}
}

func main() {

	for {
		i := r.Intn(3)

		if rgb[i] == 0 {
			for rgb[i] < 255 {
				rgb[i] = rgb[i] + 5
				print(rgb)
			}
		} else {

			for rgb[i] > 0 {
				rgb[i] = rgb[i] - 5
				print(rgb)
			}
		}
	}

}

func print(rgb [3]uint8) {
	s := fmt.Sprintf("R: %03d, G: %03d, B: %03d - #%02[1]x%02[2]x%02[3]x\n", rgb[0], rgb[1], rgb[2])

	coloredWord := rgbterm.FgString(s, rgb[0], rgb[1], rgb[2])
	fmt.Print(coloredWord)
	time.Sleep(interval)
}