package utils

import termbox "github.com/nsf/termbox-go"

// Print display a string
func Print(start Coordinates, attr Attributes, msg string) {
	x := start.X
	for _, c := range msg {
		termbox.SetCell(x, start.Y, c, attr.Foreground, attr.Background)
		x++
	}
}

// PrintWrap display a string but goes to next line if it doesn't have the place to print the string on current screen.
// It return the row number of the empty empty after the message displayed
func PrintWrap(start Coordinates, attr Attributes, msg string, margin int) int {
	w, _ := termbox.Size()
	x := start.X + margin
	y := start.Y
	for _, c := range msg {
		termbox.SetCell(x, y, c, attr.Foreground, attr.Background)
		x++

		if x >= w-margin {
			y++
			x = start.X + margin
		}
	}
	y++
	return y
}
