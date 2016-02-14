package page

import (
	"github.com/normegil/moebius-go/views/terminal/gui/utils"
	termbox "github.com/nsf/termbox-go"
)

// GoTo is a popup to go to a specific page nuber in a list
type GoTo struct {
	enabled bool
	Content string
}

// Draw function draw the popup in the terminal
func (*GoTo) Draw() error {
	w, h := termbox.Size()
	size := utils.Size{Width: 70, Height: 50}
	utils.Fill(utils.Coordinates{
		X: (w - size.Width) / 2,
		Y: (h - size.Height) / 2,
	}, size, utils.Attributes{
		Background: termbox.ColorCyan,
	})
	return nil
}

// Enabled return true if the popup is enabled
func (goTo *GoTo) Enabled() bool {
	return goTo.enabled
}
