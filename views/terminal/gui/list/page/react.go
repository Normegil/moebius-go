package page

import (
	log "github.com/Sirupsen/logrus"
	termbox "github.com/nsf/termbox-go"
)

// React listen for keyboard events and react accordingly
// Return a boolean indicating the need to redraw and a boolean indicating if the function has reacted to the event.
func (goTo *GoTo) React(ev termbox.Event) (bool, bool) {
	switch ev.Key {
	case termbox.KeyEsc:
		log.Debug("Closing GoTo Popup")
		goTo.enabled = false
		return true, true
	}
	return false, false
}
