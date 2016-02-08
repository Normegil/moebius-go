package gui

import (
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
func Launch(args views.ViewInputs) {
	err := termbox.Init()
	if nil != err {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	listerBody := &list.Lister{}
	listerBody.Init(args)

	redraw(listerBody)

eventLoop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC, termbox.KeyEsc:
				break eventLoop
			}
		case termbox.EventResize:
			redraw(listerBody)
		case termbox.EventInterrupt:
			break eventLoop
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
