package gui

import (
	log "github.com/Sirupsen/logrus"
	"github.com/normegil/moebius-go/views"
	"github.com/normegil/moebius-go/views/terminal/gui/screens/selector"
	"github.com/normegil/moebius-go/views/terminal/gui/utils"
	termbox "github.com/nsf/termbox-go"
)

var headerOpts = struct {
	titlePt int
}{
	titlePt: 5,
}

// Launch will launch the application and print a nice view of the application in terminal
func Launch(args views.ViewInputs) error {
	log.Info("Launching Termbox")
	err := termbox.Init()
	if nil != err {
		log.WithField("error", err).Fatal("Error initializing termbox")
		return err
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	log.Info("Termbox initialized")
	screen, err := selector.New(args)
	if nil != err {
		log.WithField("error", err).Fatal("Error while redrawing")
		return err
	}
	log.Info("Selector screen created")

	err = redraw(screen)
	if nil != err {
		log.WithField("error", err).Fatal("Error while drawing screen")
		return err
	}
	log.Info("Screen drawn")

eventLoop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			shouldQuit, err := reactToKey(screen, ev)
			if nil != err {
				return err
			}
			if shouldQuit {
				break eventLoop
			}
		case termbox.EventResize:
			w, h := termbox.Size()
			log.WithFields(log.Fields{
				"Width":  w,
				"Height": h,
			}).Debug("Redrawing GUI (Terminal resized)")
			needRedraw, react := screen.React(ev)
			if react && needRedraw {
				err = redraw(screen)
				if nil != err {
					return err
				}
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

func reactToKey(component utils.Component, ev termbox.Event) (bool, error) {
	if termbox.KeyCtrlC == ev.Key {
		log.Info("Exiting application (User quit)")
		return true, nil
	}

	needRedraw, treated := component.React(ev)
	if treated {
		if needRedraw {
			err := redraw(component)
			if nil != err {
				return false, err
			}
		}
		return false, nil
	}

	if termbox.KeyEsc == ev.Key {
		log.Info("Exiting application (User quit)")
		return true, nil
	}
	return false, nil
}
