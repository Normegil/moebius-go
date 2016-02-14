package list

import (
	"math"

	"github.com/normegil/moebius-go/views/terminal/gui/utils"
	termbox "github.com/nsf/termbox-go"
)

// Draw the list of mangas
func (lister *Lister) Draw(start int) error {
	w, h := termbox.Size()

	row := start
	col := 0

	endRow := h - 2

	nbCol := w / colSize
	nbRow := endRow - row
	contentSize := len(lister.content)
	nbElements := int(math.Max(1, math.Min(float64(nbCol*nbRow), float64(contentSize))))
	pageNb := (lister.selected) / nbElements

	startIndex := pageNb * nbElements
	calculatedEndIndex := (pageNb + 1) * nbElements
	endIndex := int(math.Min(float64(contentSize), float64(calculatedEndIndex)))
	for i, manga := range lister.content[startIndex:endIndex] {
		attr := utils.Attributes{Foreground: termbox.ColorWhite}
		if startIndex+i == lister.selected {
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

	lister.Footer(nbElements, len(lister.content))
	return drawPopups(lister)
}

func drawPopups(lister *Lister) error {
	return nil
}
