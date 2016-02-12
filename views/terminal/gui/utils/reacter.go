package utils

import "github.com/nsf/termbox-go"

// Reacter define a function to react to termbox events
type Reacter interface {
	React(ev termbox.Event) (bool, bool)
}
