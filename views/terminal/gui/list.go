package gui

import (
	"fmt"
	"math"
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

func (list *lister) draw(start int) error {
	w, h := termbox.Size()

	row := start
	if colSize > w {
		row = printWrap(coordinates{0, start}, attributes{foreground: termbox.ColorRed | termbox.AttrBold}, "Application won't work nicely under 40 character-wide terminal", 2)
		row++ // Leave empty space
	}

	col := 0
	endOfList := h - 2
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
			if row >= endOfList {
				break
			}
		}
	}
	list.footer()
	return nil
}

func (list *lister) footer() {
	w, h := termbox.Size()
	defaultFooterAttributes := attributes{
		foreground: termbox.ColorBlack,
		background: termbox.ColorCyan,
	}
	fill(coordinates{0, h - 1}, sizes{w, 1}, defaultFooterAttributes)

	print(coordinates{5, h - 1}, defaultFooterAttributes, fmt.Sprintf("Pages: %d/%d", 5, 100))
	mangasCountFormat := "Mangas: %d"
	sizeOfMessage := float64(len(mangasCountFormat) + 5)
	marginRight := int(math.Min(math.Max(sizeOfMessage, float64(w/8)), float64(colSize)))
	print(coordinates{w - marginRight, h - 1}, defaultFooterAttributes, fmt.Sprintf(mangasCountFormat, len(list.content)))
}
