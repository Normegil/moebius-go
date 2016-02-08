package gui

import (
	"github.com/normegil/moebius-go/views/terminal/gui/utils"
	termbox "github.com/nsf/termbox-go"
)

func drawHeader() {
	w, _ := termbox.Size()
	utils.Fill(utils.Coordinates{0, 0}, utils.Size{w, 1}, utils.Attributes{
		Foreground: termbox.ColorBlack,
		Background: termbox.ColorCyan,
	})
	utils.Print(utils.Coordinates{headerOpts.titlePt, 0}, utils.Attributes{
		Foreground: termbox.ColorBlack | termbox.AttrBold,
		Background: termbox.ColorCyan,
	}, "Moebius-go")
}
