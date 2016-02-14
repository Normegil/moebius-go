package list

import (
	log "github.com/Sirupsen/logrus"
	termbox "github.com/nsf/termbox-go"
)

// MoveUp move the selected element up on the screen
// Returns true if a redraw is needed
func (lister *Lister) MoveUp() bool {
	log.Debug("Moving Up")
	w, _ := termbox.Size()
	numberOfElementsByColumn := w / colSize
	newIndex := lister.selected - numberOfElementsByColumn
	if 0 > newIndex {
		log.Debug("Already at highest row")
		return false
	}
	lister.selected = newIndex
	return true
}

// MoveDown move the selected element down on the screen
// Returns true if a redraw is needed
func (lister *Lister) MoveDown() bool {
	log.Debug("Moving Down")
	w, _ := termbox.Size()
	numberOfElementsByColumn := w / colSize
	newIndex := lister.selected + numberOfElementsByColumn
	if len(lister.content) < newIndex {
		log.Debug("Already at lowest row")
		return false
	}
	lister.selected = newIndex
	return true
}

// MoveLeft move the selected element left on the screen
// Returns true if a redraw is needed
func (lister *Lister) MoveLeft() bool {
	log.Debug("Moving Left")
	w, _ := termbox.Size()
	numberOfElementsByColumn := w / colSize
	indexStartingFrom1 := lister.selected + 1
	if indexStartingFrom1%numberOfElementsByColumn == 1 {
		log.Debug("No element at left side")
		return false
	}
	lister.selected--
	return true
}

// MoveRight move the selected element up on the screen
// Returns true if a redraw is needed
func (lister *Lister) MoveRight() bool {
	log.Debug("Moving Right")
	w, _ := termbox.Size()
	numberOfElementsByColumn := w / colSize
	indexStartingFrom1 := lister.selected + 1
	if indexStartingFrom1%numberOfElementsByColumn == 0 {
		log.Debug("No element at right side")
		return false
	}
	lister.selected++
	return true
}
