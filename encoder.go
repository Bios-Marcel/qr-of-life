package main

import (
	"fmt"
	"image"
	"io"
	"strconv"
	"strings"
)

type Encoder struct {
}

func (e Encoder) Encode(w io.Writer, img image.Image) error {
	imgBounds := img.Bounds()
	width := imgBounds.Dx()
	height := imgBounds.Dy()

	var data strings.Builder
	for y := 0; y < height; y++ {
		data.WriteString(fmt.Sprintf("{\"%d\":[", y+27))
		var blackIndices []string
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if r == 0 && g == 0 && b == 0 {
				blackIndices = append(blackIndices, strconv.FormatInt(int64(x)+76, 10))
			}
		}
		data.WriteString(strings.Join(blackIndices, ","))
		data.WriteString("]},")
	}

	_, err := fmt.Fprintf(w, "https://pmav.eu/stuff/javascript-game-of-life-v3.1.1/?autoplay=0&trail=1&grid=1&colors=1&zoom=1&s=[%s]\n", data.String())
	return err
}
