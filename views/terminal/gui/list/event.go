package list

import "github.com/nsf/termbox-go"

// React listen for keyboard events and react accordingly
// Return a boolean indicating the need to redraw
func (lister *Lister) React(ev termbox.Event) bool {
	switch ev.Key {
	case termbox.KeyArrowUp:
		return lister.MoveUp()
	case termbox.KeyArrowDown:
		return lister.MoveDown()
	case termbox.KeyArrowLeft:
		return lister.MoveLeft()
	case termbox.KeyArrowRight:
		return lister.MoveRight()
	}
	return false
}
