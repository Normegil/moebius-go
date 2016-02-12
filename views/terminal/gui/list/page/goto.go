package page

import (
	"github.com/normegil/moebius-go/views/terminal/gui/utils"
	"github.com/nsf/termbox-go"
)

type GoTo struct {
	enabled bool
	Content string
}

func (gotTo *GoTo) Draw() error {
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

func (goTo *GoTo) Enabled() bool {
	return goTo.Enabled()
}
