package search

import (
	"math"

	"github.com/normegil/moebius-go/views/terminal/gui/utils"
	termbox "github.com/nsf/termbox-go"
)

const minWidth = 40
const col = 10

// Draw draws the search field
func (field *Search) Draw(start int) error {

	field.active = true
	row := start
	attributes := utils.Attributes{}
	if field.active {
		attributes = utils.Attributes{Foreground: termbox.AttrReverse, Background: termbox.AttrReverse}
		w, _ := termbox.Size()
		width := int(math.Max(float64(minWidth), float64(w-col*2)))
		utils.Fill(utils.Coordinates{X: col, Y: row}, utils.Size{Width: width, Height: 1}, attributes)
	}

	utils.Print(utils.Coordinates{X: col, Y: row}, attributes, "> Test"+field.content)
	return nil
}
