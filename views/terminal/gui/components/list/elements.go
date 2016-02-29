package list

import (
	"math"

	log "github.com/Sirupsen/logrus"
)

// Displayed get the number of element displayed currently on screen
func (list List) Displayed() int {
	nbCol := list.getNumberOfColumns()
	nbRow := list.getNumberOfRows()
	contentSize := len(list.Displayable)

	displayed := int(math.Min(float64(nbCol*nbRow), float64(contentSize)))
	log.WithFields(log.Fields{
		"Start":        list.Start,
		"End":          list.End,
		"Columns":      nbCol,
		"Rows":         nbRow,
		"Content size": contentSize,
		"Displayed":    displayed,
	}).Debug("Display size calculated")

	return displayed
}

func (list List) getNumberOfColumns() int {
	return (list.End.X - list.Start.X) / list.CellSize.Width
}

func (list List) getNumberOfRows() int {
	return ((list.End.Y - 1) - list.Start.Y) / list.CellSize.Height
}
