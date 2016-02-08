package gui

import (
	"sort"

	"github.com/normegil/moebius-go/connector"
	"github.com/normegil/moebius-go/views"
	termbox "github.com/nsf/termbox-go"
)

const colSize = 40
const printSize = 30
const marginLeft = (colSize - printSize) / 2
const overflowPrintSize = printSize - 3

type lister struct {
	content  []string
	selected string
}

func (list *lister) init(args views.ViewInputs) error {
	mangas, err := connector.ListMangas(args.Cache, args.Listers, connector.ListMangasOptions{
		UseCache: true,
		Language: "en",
	})
	if nil != err {
		return err
	}
	list.content = make([]string, len(mangas))
	for i, manga := range mangas {
		toPrint := manga.Title
		if len(toPrint) > printSize {
			toPrint = toPrint[0:overflowPrintSize] + "..."
		}
		list.content[i] = toPrint
	}
	sort.Strings(list.content)
	list.selected = list.content[0]
	return nil
}

func (list *lister) draw() error {
	w, _ := termbox.Size()

	row := 2
	col := 0
	for _, toPrint := range list.content {
		attr := attributes{foreground: termbox.ColorWhite}
		if toPrint == list.selected {
			attr = attributes{foreground: termbox.AttrReverse, background: termbox.AttrReverse}
		}
		print(coordinates{col + marginLeft, row}, attr, toPrint)

		col += colSize
		if col+colSize >= w {
			col = 0
			row++
		}
	}
	return nil
}
