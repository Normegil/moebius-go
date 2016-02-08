package gui

import termbox "github.com/nsf/termbox-go"

func redraw(toDraw body) error {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)

	drawHeader()
	toDraw.Draw(2)

	err := termbox.Flush()
	return err
}
