package selector

import (
	"github.com/normegil/moebius-go/views/terminal/gui/utils"
	termbox "github.com/nsf/termbox-go"
)

// React listen for keyboard events and react accordingly
// Return a boolean indicating the need to redraw and a boolean indicating if the function has reacted to the event.
func (selector *Selector) React(ev termbox.Event) (bool, bool) {

	if termbox.EventResize == ev.Type {
		w, h := termbox.Size()
		selector.list.End = utils.Coordinates{X: w, Y: h - 1}
		return true, true
	}

	switch ev.Key {
	case termbox.KeyTab:
		selector.switchActive()
		return true, true
	default:
		redraw, react := subComponentReact(selector, ev)
		if react {
			return redraw, true
		}
	}
	return false, false
}

func (selector *Selector) switchActive() {
	if selector.active == searchID {
		selector.active = listID
	} else if selector.active == listID {
		selector.active = searchID
	}
}

func subComponentReact(selector *Selector, event termbox.Event) (bool, bool) {
	if selector.active == searchID {
		needRedraw, hasReacted := selector.search.React(event)
		newFilter := string(selector.search.Content)
		selector.list.Filter(newFilter)
		return needRedraw, hasReacted
	} else if selector.active == listID {
		return selector.list.React(event)
	}
	return false, false
}
