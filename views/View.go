package views

import (
	cachePkg "github.com/normegil/moebius-go/cache"
	"github.com/normegil/moebius-go/connector"
)

// View represent a graphical or terminal view
type View interface {
	Launch(inputs ViewInputs) error
}

// ViewInputs gather the input needed by guis to launch the application
type ViewInputs struct {
	Cache   cachePkg.Cache
	Listers []connector.Lister
}
