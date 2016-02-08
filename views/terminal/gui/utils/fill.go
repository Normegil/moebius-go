package utils

import termbox "github.com/nsf/termbox-go"

// Fill create a rectangle of given dimension, with the upper left corned defined by Coordinates. It fills the cells with defined Attributes
func Fill(start Coordinates, s Size, attr Attributes) {
	blockWidth := s.Width - start.X
	for i := 0; i < blockWidth; i++ {
		for j := 0; j < s.Height; j++ {
			termbox.SetCell(start.X+i, start.Y+j, ' ', attr.Foreground, attr.Background)
		}
	}
}
