package gui

import termbox "github.com/nsf/termbox-go"

func redraw(toDraw body) error {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)

	_, h := termbox.Size()
	drawHeader()
	toDraw.draw(2, h-2)

	err := termbox.Flush()
	return err
}
