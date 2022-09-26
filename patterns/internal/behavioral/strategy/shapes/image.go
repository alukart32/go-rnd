package shapes

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"

	"alukart32.com/usage/patterns/internal/behavioral/strategy"
)

type ImageSquare struct {
	strategy.PrintOutput
}

func (t *ImageSquare) Print() error {
	if t.Writer == nil {
		return fmt.Errorf("no writer stored on ImageSquare")
	}

	width := 800
	height := 600
	origin := image.Point{0, 0}

	// background for pic
	bgImage := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})

	// gray background color
	bgColor := image.Uniform{color.RGBA{R: 70, G: 70, B: 70, A: 0}}

	// draw [where; in what bounds; color; bounds start; operation type]
	draw.Draw(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)

	// pic
	squareWidth := 200
	squareHeight := squareWidth
	squareColor := image.Uniform{color.RGBA{R: 255, G: 0, B: 0, A: 1}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	// move pic to the center
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	squareImage := image.NewRGBA(square)

	// draw a pic to background
	draw.Draw(bgImage, squareImage.Bounds(), &squareColor, origin, draw.Src)

	// jpeg.Encode the image.
	quality := &jpeg.Options{Quality: 75}
	if err := jpeg.Encode(t.Writer, bgImage, quality); err != nil {
		return fmt.Errorf("error writing image to disk")
	}

	if t.LogWriter != nil {
		t.LogWriter.Write([]byte("Image written in provided writer\n"))
	}
	return nil
}
