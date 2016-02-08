package list

import "github.com/nsf/termbox-go"

// MoveUp move the selected element up on the screen
// Returns true if a redraw is needed
func (lister *Lister) MoveUp() bool {
	w, _ := termbox.Size()
	numberOfElementsByColumn := w / colSize
	newIndex := lister.selected - numberOfElementsByColumn
	if 0 > newIndex {
		return false
	}
	lister.selected = newIndex
	return true
}

// MoveDown move the selected element down on the screen
// Returns true if a redraw is needed
func (lister *Lister) MoveDown() bool {
	w, _ := termbox.Size()
	numberOfElementsByColumn := w / colSize
	newIndex := lister.selected + numberOfElementsByColumn
	if 0 > newIndex {
		return false
	}
	lister.selected = newIndex
	return true
}

// MoveLeft move the selected element left on the screen
// Returns true if a redraw is needed
func (lister *Lister) MoveLeft() bool {
	w, _ := termbox.Size()
	numberOfElementsByColumn := w / colSize
	indexStartingFrom1 := lister.selected + 1
	if indexStartingFrom1%numberOfElementsByColumn == 1 {
		return false
	}
	lister.selected--
	return true
}

// MoveRight move the selected element up on the screen
// Returns true if a redraw is needed
func (lister *Lister) MoveRight() bool {
	w, _ := termbox.Size()
	numberOfElementsByColumn := w / colSize
	indexStartingFrom1 := lister.selected + 1
	if indexStartingFrom1%numberOfElementsByColumn == 0 {
		return false
	}
	lister.selected++
	return true
}
