package input

import (
	log "github.com/Sirupsen/logrus"
	termbox "github.com/nsf/termbox-go"
)

// React manage event happening on the input field
func (input *Input) React(ev termbox.Event) (bool, bool) {
	log.WithFields(log.Fields{
		"Key": ev.Key,
		"Ch":  string(ev.Ch),
	}).Info("Event on input field")

	if ev.Ch != 0 {
		insertRune(input, ev.Ch)
		input.CursorPosition++
		return true, true
	} else if ev.Key != 0 {
		return input.reactToKey(ev)
	}
	return false, false
}

func (input *Input) reactToKey(ev termbox.Event) (bool, bool) {
	if (termbox.KeyBackspace == ev.Key || termbox.KeyBackspace2 == ev.Key) && len(input.Content) != 0 && input.CursorPosition != 0 {
		input.Content = removeBytesFrom(input.Content, input.CursorPosition-1, input.CursorPosition)
		input.CursorPosition--
		return true, true
	} else if termbox.KeyDelete == ev.Key && len(input.Content) != 0 && input.CursorPosition < len(input.Content) {
		input.Content = removeBytesFrom(input.Content, input.CursorPosition, input.CursorPosition+1)
		return true, true
	} else if termbox.KeyArrowLeft == ev.Key && input.CursorPosition > 0 {
		input.CursorPosition--
		return true, true
	} else if termbox.KeyArrowRight == ev.Key && input.CursorPosition < len(input.Content) {
		input.CursorPosition++
		return true, true
	} else if termbox.KeyHome == ev.Key {
		input.CursorPosition = 0
		return true, true
	} else if termbox.KeyEnd == ev.Key {
		input.CursorPosition = len(input.Content)
		return true, true
	} else if termbox.KeySpace == ev.Key {
		insertRune(input, ' ')
		input.CursorPosition++
		return true, true
	}
	return false, false
}
