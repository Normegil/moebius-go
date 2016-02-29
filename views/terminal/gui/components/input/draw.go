package input

import (
	log "github.com/Sirupsen/logrus"
	"github.com/normegil/moebius-go/views/terminal/gui/utils"
	termbox "github.com/nsf/termbox-go"
)

const minWidth = 40

// Draw draws the search field
func (input *Input) Draw(active bool) error {
	attr := getAttributes(active)
	input.print(attr, input.Prefix+string(input.Content))

	hOffset := input.Start.X + 2
	termbox.SetCursor(input.CursorPosition+hOffset, input.Start.Y)
	if !active {
		termbox.HideCursor()
	}

	return nil
}

func getAttributes(active bool) utils.Attributes {
	if active {
		return utils.Attributes{Foreground: termbox.AttrReverse, Background: termbox.AttrReverse}
	}
	return utils.Attributes{}
}

func (input *Input) print(attributes utils.Attributes, toPrint string) {
	log.WithFields(log.Fields{
		"start": input.Start.X,
		"end":   input.End,
	}).Debug("Print input string")
	utils.Fill(utils.Coordinates{X: input.Start.X, Y: input.Start.Y}, utils.Size{Width: input.End - input.Start.X, Height: 1}, attributes)
	utils.Print(utils.Coordinates{X: input.Start.X, Y: input.Start.Y}, attributes, toPrint)
}
