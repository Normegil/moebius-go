package gui

import (
	log "github.com/Sirupsen/logrus"
	termbox "github.com/nsf/termbox-go"
)

func redraw(toDraw body) error {
	const coldef = termbox.ColorDefault
	err := termbox.Clear(coldef, coldef)
	if nil != err {
		log.WithField("error", err).Warn("Error while clearing terminal")
		return err
	}

	drawHeader()
	err = toDraw.Draw(2)
	if nil != err {
		log.WithField("error", err).Warn("Error while drawing application")
		return err
	}

	err = termbox.Flush()
	if nil != err {
		log.WithField("error", err).Warn("Error while flushing application GUI")
	}
	return err
}
