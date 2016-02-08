package list

import (
	"math"

	"github.com/normegil/moebius-go/views/terminal/gui/utils"
	"github.com/nsf/termbox-go"
)

// Draw the list of mangas
func (list *Lister) Draw(start int) error {
	w, h := termbox.Size()

	row := start
	if colSize > w {
		row = utils.PrintWrap(utils.Coordinates{X: 0, Y: start}, utils.Attributes{
			Foreground: termbox.ColorRed | termbox.AttrBold,
		}, "Application won't work nicely under 40 character-wide terminal", 2)
		row++ // Leave empty space
	}

	col := 0
	endOfList := h - 2
	for i, manga := range list.content {
		attr := utils.Attributes{Foreground: termbox.ColorWhite}
		if i == list.selected {
			attr = utils.Attributes{
				Foreground: termbox.AttrReverse,
				Background: termbox.AttrReverse,
			}
		}
		utils.Print(utils.Coordinates{X: col + marginLeft, Y: row}, attr, format(manga))

		col += colSize
		if col+colSize >= w {
			col = 0
			row++
			if row >= endOfList {
				break
			}
		}
	}

	nbCol := w / colSize
	nbRow := endOfList - start
	elementsDisplayed := int(math.Max(1, math.Min(float64(nbCol*nbRow), float64(len(list.content)))))
	list.Footer(elementsDisplayed, len(list.content))
	return nil
}
