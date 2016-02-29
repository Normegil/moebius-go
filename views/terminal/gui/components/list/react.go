package list

import termbox "github.com/nsf/termbox-go"

// React listen for keyboard events and react accordingly
// Return a boolean indicating the need to redraw and a boolean indicating if the function has reacted to the event.
func (list *List) React(ev termbox.Event) (bool, bool) {
	switch ev.Key {
	case termbox.KeyArrowUp:
		return list.MoveUp(), true
	case termbox.KeyArrowDown:
		return list.MoveDown(), true
	case termbox.KeyArrowLeft:
		return list.MoveLeft(), true
	case termbox.KeyArrowRight:
		return list.MoveRight(), true
	case termbox.KeyHome:
		return list.MoveStart(), true
	case termbox.KeyEnd:
		return list.MoveEnd(), true
	case termbox.KeyPgup:
		return list.MovePageUp(), true
	case termbox.KeyPgdn:
		return list.MovePageDown(), true
	}
	return false, false
}
