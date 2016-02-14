package list

import termbox "github.com/nsf/termbox-go"

// React listen for keyboard events and react accordingly
// Return a boolean indicating the need to redraw and a boolean indicating if the function has reacted to the event.
func (lister *Lister) React(ev termbox.Event) (bool, bool) {

	switch ev.Key {
	case termbox.KeyArrowUp:
		return lister.MoveUp(), true
	case termbox.KeyArrowDown:
		return lister.MoveDown(), true
	case termbox.KeyArrowLeft:
		return lister.MoveLeft(), true
	case termbox.KeyArrowRight:
		return lister.MoveRight(), true
	}
	return false, false
}
