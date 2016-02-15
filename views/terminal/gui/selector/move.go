package list

import (
	"math"

	log "github.com/Sirupsen/logrus"
	termbox "github.com/nsf/termbox-go"
)

// MoveUp move the selected element up on the screen
// Returns true if a redraw is needed
func (list *Lister) MoveUp() bool {
	log.Debug("Moving Up")
	w, _ := termbox.Size()
	numberOfElementsByColumn := w / colSize
	newIndex := list.selected - numberOfElementsByColumn
	if 0 > newIndex {
		log.Debug("Already at highest row")
		return false
	}
	list.selected = newIndex
	return true
}

// MoveDown move the selected element down on the screen
// Returns true if a redraw is needed
func (list *Lister) MoveDown() bool {
	log.Debug("Moving Down")
	w, _ := termbox.Size()
	numberOfElementsByColumn := w / colSize
	newIndex := list.selected + numberOfElementsByColumn
	if len(list.content)-1 < newIndex {
		log.Debug("Already at lowest row")
		return false
	}
	list.selected = newIndex
	return true
}

// MoveLeft move the selected element left on the screen
// Returns true if a redraw is needed
func (list *Lister) MoveLeft() bool {
	log.Debug("Moving Left")
	w, _ := termbox.Size()
	numberOfElementsByColumn := w / colSize
	indexStartingFrom1 := list.selected + 1
	if indexStartingFrom1%numberOfElementsByColumn == 1 {
		log.Debug("No element at left side")
		return false
	}
	list.selected--
	return true
}

// MoveRight move the selected element up on the screen
// Returns true if a redraw is needed
func (list *Lister) MoveRight() bool {
	log.Debug("Moving Right")
	w, _ := termbox.Size()
	numberOfElementsByColumn := w / colSize
	indexStartingFrom1 := list.selected + 1
	if indexStartingFrom1%numberOfElementsByColumn == 0 {
		log.Debug("No element at right side")
		return false
	}
	list.selected++
	return true
}

// MoveStart move the cusor at the start of the displayed list
func (list *Lister) MoveStart() bool {
	if 0 == list.selected {
		return false
	}
	log.Debug("Moving at start of the content")
	list.selected = 0
	return true
}

// MoveEnd move the cusor at the end of the displayed list
func (list *Lister) MoveEnd() bool {
	nbElements := len(list.content)
	if nbElements-1 == list.selected {
		return false
	}
	list.selected = nbElements - 1
	log.WithFields(log.Fields{
		"content size":    nbElements,
		"cursor position": list.selected,
	}).Debug("Moving at end of the content")
	return true
}

// MovePageUp move one page down if it's possible
func (list *Lister) MovePageUp() bool {
	index := list.selected
	contentSize := len(list.content)
	index -= getNbElementPerPage(contentSize)
	list.selected = int(math.Max(0, float64(index)))
	return true
}

// MovePageDown move one page down if it's possible
func (list *Lister) MovePageDown() bool {
	index := list.selected
	contentSize := len(list.content)
	index += getNbElementPerPage(contentSize)
	list.selected = int(math.Min(float64(contentSize-1), float64(index)))
	return true
}

func getNbElementPerPage(contentSize int) int {
	w, h := termbox.Size()
	row := 2
	endRow := h - 2
	nbCol := w / colSize
	nbRow := endRow - row
	return int(math.Max(1, math.Min(float64(nbCol*nbRow), float64(contentSize))))
}
