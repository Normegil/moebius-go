package gui

import (
	log "github.com/Sirupsen/logrus"
	"github.com/normegil/moebius-go/views"
	"github.com/normegil/moebius-go/views/terminal/gui/list"
	termbox "github.com/nsf/termbox-go"
)

var headerOpts = struct {
	titlePt int
}{
	titlePt: 5,
}

// Launch will launch the application and print a nice view of the application in terminal
func Launch(args views.ViewInputs) error {
	err := termbox.Init()
	if nil != err {
		return err
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	lister := &list.Lister{}
	err = lister.Init(args)
	if nil != err {
		return err
	}

	err = redraw(lister)
	if nil != err {
		return err
	}

eventLoop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC:
				log.Info("Exiting application (User quit)")
				break eventLoop
			default:
				needRedraw, treated := lister.React(ev)
				if treated {
					if needRedraw {
						err = redraw(lister)
						if nil != err {
							return err
						}
					}
				} else if ev.Key == termbox.KeyEsc {
					log.Info("Exiting application (User quit)")
					break eventLoop
				}
			}
		case termbox.EventResize:
			log.Debug("Redrawing GUI (Terminal resized)")
			err = redraw(lister)
			if nil != err {
				return err
			}
		case termbox.EventInterrupt:
			log.Info("Interrupting application (SIGINT received))")
			break eventLoop
		case termbox.EventError:
			return ev.Err
		}
	}
	return nil
}
