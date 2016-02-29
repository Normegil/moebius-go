package input

import "github.com/normegil/moebius-go/views/terminal/gui/utils"

// Input is an editable field
type Input struct {
	Prefix         string
	Start          utils.Coordinates
	End            int
	Content        []byte
	CursorPosition int
}
