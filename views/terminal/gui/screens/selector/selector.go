package selector

import (
	"github.com/normegil/moebius-go/views/terminal/gui/components/input"
	"github.com/normegil/moebius-go/views/terminal/gui/components/list"
)

const searchID = 0
const listID = 1

// Selector manage the screen listing mangas
type Selector struct {
	search *input.Input
	list   *list.List
	active int
}
