package selector

import (
	"fmt"
	"sort"

	"github.com/normegil/moebius-go/views/terminal/gui/components/input"
	"github.com/normegil/moebius-go/views/terminal/gui/components/list"

	termbox "github.com/nsf/termbox-go"

	"github.com/normegil/moebius-go/connector"
	"github.com/normegil/moebius-go/models"
	"github.com/normegil/moebius-go/views"
	"github.com/normegil/moebius-go/views/terminal/gui/utils"
)

// New initialize and load needed data to display the selector
func New(args views.ViewInputs) (*Selector, error) {
	content, err := getContent(args)
	if nil != err {
		return nil, err
	}

	w, h := termbox.Size()
	selector := Selector{
		list: &list.List{
			Start:           utils.Coordinates{X: 4, Y: 4},
			End:             utils.Coordinates{X: w, Y: h - 1},
			CellSize:        utils.Size{Width: 40, Height: 1},
			PrintSize:       40 - 5,
			OriginalContent: content,
		},
		search: &input.Input{
			Prefix:         "> ",
			Start:          utils.Coordinates{X: 4, Y: 2},
			End:            w,
			Content:        []byte{},
			CursorPosition: 0,
		},
		active: searchID,
	}
	selector.list.Filter("")
	return &selector, nil
}

func getContent(args views.ViewInputs) ([]fmt.Stringer, error) {
	mangas, err := connector.ListMangas(args.Cache, args.Listers, connector.ListMangasOptions{
		UseCache: true,
		Language: "en",
	})
	if nil != err {
		return nil, err
	}
	sort.Sort(models.ByTitle(mangas))
	var stringers []fmt.Stringer
	for _, manga := range mangas {
		stringers = append(stringers, manga)
	}
	return stringers, nil
}
