package list

import (
	"math"

	log "github.com/Sirupsen/logrus"
)

// MoveUp move the selected element up on the screen
// Returns true if a redraw is needed
func (list *List) MoveUp() bool {
	log.WithField("selected", list.Selected).Debug("Moving Up")
	newIndex := list.Selected - list.getNumberOfColumns()
	if 0 > newIndex {
		log.Debug("Already at highest row")
		return false
	}
	list.Selected = newIndex
	return true
}

// MoveDown move the selected element down on the screen
// Returns true if a redraw is needed
func (list *List) MoveDown() bool {
	log.WithField("selected", list.Selected).Debug("Moving Down")
	newIndex := list.Selected + list.getNumberOfColumns()
	lastPossibleIndex := len(list.Displayable) - 1
	if lastPossibleIndex < newIndex {
		log.Debug("Already at lowest row")
		return false
	}
	list.Selected = newIndex
	return true
}

// MoveLeft move the selected element left on the screen
// Returns true if a redraw is needed
func (list *List) MoveLeft() bool {
	log.WithField("selected", list.Selected).Debug("Moving Left")
	indexStartingFrom1 := list.Selected + 1
	if indexStartingFrom1%list.getNumberOfColumns() == 1 || list.Selected == 0 {
		log.Debug("No element at left side")
		return false
	}
	list.Selected--
	return true
}

// MoveRight move the selected element up on the screen
// Returns true if a redraw is needed
func (list *List) MoveRight() bool {
	log.WithField("selected", list.Selected).Debug("Moving Right")
	indexStartingFrom1 := list.Selected + 1
	lastPossibleIndex := len(list.Displayable) - 1
	if indexStartingFrom1%list.getNumberOfColumns() == 0 || list.Selected >= lastPossibleIndex {
		log.Debug("No element at right side")
		return false
	}
	list.Selected++
	return true
}

// MoveStart move the cusor at the start of the displayed list
func (list *List) MoveStart() bool {
	if 0 == list.Selected {
		return false
	}
	log.Debug("Moving at start of the content")
	list.Selected = 0
	return true
}

// MoveEnd move the cusor at the end of the displayed list
func (list *List) MoveEnd() bool {
	list.Selected = len(list.Displayable) - 1
	return true
}

// MovePageUp move one page down if it's possible
func (list *List) MovePageUp() bool {
	index := list.Selected - list.Displayed()
	list.Selected = int(math.Max(0, float64(index)))
	return true
}

// MovePageDown move one page down if it's possible
func (list *List) MovePageDown() bool {
	contentSize := len(list.Displayable)
	index := list.Selected + list.Displayed()
	list.Selected = int(math.Min(float64(contentSize-1), float64(index)))
	return true
}
