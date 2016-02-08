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
	col := 0

	endRow := h - 2

	nbCol := w / colSize
	nbRow := endRow - row
	nbElements := int(math.Max(1, math.Min(float64(nbCol*nbRow), float64(len(list.content)))))
	pageNb := (list.selected) / nbElements

	startIndex := pageNb * nbElements
	endIndex := (pageNb + 1) * nbElements
	for i, manga := range list.content[startIndex:endIndex] {
		attr := utils.Attributes{Foreground: termbox.ColorWhite}
		if startIndex+i == list.selected {
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
			if row >= endRow {
				break
			}
		}
	}

	list.Footer(nbElements, len(list.content))
	return nil
}
