package list

import (
	"fmt"

	"github.com/normegil/moebius-go/views/terminal/gui/utils"
)

// List display a list of selectable elements
type List struct {
	Start, End      utils.Coordinates
	CellSize        utils.Size
	PrintSize       int
	Selected        int
	OriginalContent []fmt.Stringer
	Displayable     []fmt.Stringer
}
