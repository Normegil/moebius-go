package list

import (
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
	"github.com/normegil/moebius-go/views/terminal/gui/utils"
	termbox "github.com/nsf/termbox-go"
)

// Draw the list of elements
func (list *List) Draw(active bool) error {
	col := 0
	row := 0
	startIndex := list.getStartIndex()
	for i, toDisplay := range getSliceToDisplay(*list) {
		attributes := getAttributes(*list, active, startIndex+i)

		utils.Print(utils.Coordinates{X: list.Start.X + col, Y: list.Start.Y + row}, attributes, format(toDisplay.String(), list.PrintSize))

		col += list.CellSize.Width
		if !canFitCell(list.Start.X+col, list.CellSize.Width, list.End.X) {
			col = 0
			row += list.CellSize.Height
			if endPageReached(list.Start.Y+row, list.CellSize.Height, list.End.Y) {
				break
			}
		}
	}

	return nil
}

func canFitCell(column, cellWidth, endIndex int) bool {
	return column+cellWidth < endIndex
}

func endPageReached(row, cellHeight, endIndex int) bool {
	return row+cellHeight >= endIndex
}

func (list List) getStartIndex() int {
	nbElements := list.Displayed()
	pageNb := list.Page()
	return pageNb * nbElements
}

func getSliceToDisplay(list List) []fmt.Stringer {
	startIndex := list.getStartIndex()

	nbElements := list.Displayed()
	pageNb := list.Page()
	contentSize := len(list.Displayable)

	calculatedEndIndex := (pageNb + 1) * nbElements
	endIndex := int(math.Min(float64(contentSize), float64(calculatedEndIndex)))

	log.WithFields(log.Fields{
		"number of displayed elements": nbElements,
		"pag number":                   pageNb,
		"selected":                     list.Selected,
		"start":                        startIndex,
		"end":                          endIndex,
		"content size":                 len(list.Displayable),
	}).Debug("Filter slice for part to display")
	return list.Displayable[startIndex:endIndex]
}

func getAttributes(list List, active bool, index int) utils.Attributes {
	attr := utils.Attributes{Foreground: termbox.ColorWhite}
	if index == list.Selected && active {
		attr = utils.Attributes{
			Foreground: termbox.AttrReverse,
			Background: termbox.AttrReverse,
		}
	}
	return attr
}
