package gui

import termbox "github.com/nsf/termbox-go"

type attributes struct {
	foreground termbox.Attribute
	background termbox.Attribute
}

type coordinates struct {
	x int
	y int
}

type sizes struct {
	width  int
	height int
}

func print(start coordinates, attr attributes, msg string) {
	x := start.x
	for _, c := range msg {
		termbox.SetCell(x, start.y, c, attr.foreground, attr.background)
		x++
	}
}

func fill(start coordinates, s sizes, attr attributes) {
	blockWidth := s.width - start.x
	for i := 0; i < blockWidth; i++ {
		for j := 0; j < s.height; j++ {
			termbox.SetCell(start.x+i, start.y+j, ' ', attr.foreground, attr.background)
		}
	}
}
