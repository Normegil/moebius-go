package list

import (
	log "github.com/Sirupsen/logrus"
	"github.com/normegil/moebius-go/views/terminal/gui/list/page"
	termbox "github.com/nsf/termbox-go"
)

// React listen for keyboard events and react accordingly
// Return a boolean indicating the need to redraw and a boolean indicating if the function has reacted to the event.
func (lister *Lister) React(ev termbox.Event) (bool, bool) {

	if nil != lister.goTo && lister.goTo.Enabled() {
		redraw, treated := lister.goTo.React(ev)
		if treated {
			return redraw, treated
		}
	}

	switch ev.Key {
	case termbox.KeyArrowUp:
		return lister.MoveUp(), true
	case termbox.KeyArrowDown:
		return lister.MoveDown(), true
	case termbox.KeyArrowLeft:
		return lister.MoveLeft(), true
	case termbox.KeyArrowRight:
		return lister.MoveRight(), true
	case termbox.KeyCtrlP:
		openGoToPopup(lister)
		return true, true
	}
	return false, false
}

func openGoToPopup(lister *Lister) {
	if nil == lister.goTo {
		log.Debug("Init GoTo popup")
		lister.goTo = &page.GoTo{}
	}
	log.Info("Open GoTo popup")
	lister.goTo.Open("1")
}
