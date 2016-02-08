package gui

import termbox "github.com/nsf/termbox-go"

func redraw(toDraw body) error {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)

	drawHeader()
	toDraw.draw()

	err := termbox.Flush()
	return err
}
