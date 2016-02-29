package gui

import (
	log "github.com/Sirupsen/logrus"
	"github.com/normegil/moebius-go/views/terminal/gui/utils"
	termbox "github.com/nsf/termbox-go"
)

func redraw(toDraw utils.Drawer) error {
	log.Debug("Redrawing screen")
	const coldef = termbox.ColorDefault
	err := termbox.Clear(coldef, coldef)
	if nil != err {
		log.WithField("error", err).Warn("Error while clearing terminal")
		return err
	}
	log.Debug("Screen cleared")

	drawHeader()
	log.Debug("Header drawn")
	err = toDraw.Draw(true)
	if nil != err {
		log.WithField("error", err).Warn("Error while drawing application")
		return err
	}
	log.Debug("Component drawn")

	err = termbox.Flush()
	if nil != err {
		log.WithField("error", err).Warn("Error while flushing application GUI")
	}
	log.Debug("Flush done")
	return err
}
