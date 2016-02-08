package list

import (
	"fmt"
	"math"
	"sort"

	"github.com/normegil/moebius-go/connector"
	"github.com/normegil/moebius-go/views"
	"github.com/normegil/moebius-go/views/terminal/gui/utils"
	termbox "github.com/nsf/termbox-go"
)

const colSize = 40
const printSize = 30
const marginLeft = (colSize - printSize) / 2
const overflowPrintSize = printSize - 3

// Lister manage the display of a list of manga
type Lister struct {
	content  []string
	selected int
}

// Init initialize and load needed data into Lister
func (list *Lister) Init(args views.ViewInputs) error {
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
	list.selected = 163
	return nil
}

// Draw the list of mangas
func (list *Lister) Draw(start int) error {
	w, h := termbox.Size()

	row := start
	if colSize > w {
		row = utils.PrintWrap(utils.Coordinates{X: 0, Y: start}, utils.Attributes{
			Foreground: termbox.ColorRed | termbox.AttrBold,
		}, "Application won't work nicely under 40 character-wide terminal", 2)
		row++ // Leave empty space
	}

	col := 0
	endOfList := h - 2
	for i, toPrint := range list.content {
		attr := utils.Attributes{Foreground: termbox.ColorWhite}
		if i == list.selected {
			attr = utils.Attributes{
				Foreground: termbox.AttrReverse,
				Background: termbox.AttrReverse,
			}
		}
		utils.Print(utils.Coordinates{X: col + marginLeft, Y: row}, attr, toPrint)

		col += colSize
		if col+colSize >= w {
			col = 0
			row++
			if row >= endOfList {
				break
			}
		}
	}

	nbCol := w / colSize
	nbRow := endOfList - start
	elementsDisplayed := int(math.Min(float64(nbCol*nbRow), float64(len(list.content))))
	list.Footer(elementsDisplayed, len(list.content))
	return nil
}

//Footer manage the display of the footer under the list of mangas
func (list *Lister) Footer(displayed, total int) {
	w, h := termbox.Size()
	defaultFooterAttributes := utils.Attributes{
		Foreground: termbox.ColorBlack,
		Background: termbox.ColorCyan,
	}
	utils.Fill(utils.Coordinates{X: 0, Y: h - 1}, utils.Size{Width: w, Height: 1}, defaultFooterAttributes)

	elementSelectedStartingAt1 := list.selected + 1
	pageNb := elementSelectedStartingAt1 / displayed
	if elementSelectedStartingAt1%displayed != 0 {
		pageNb++
	}
	nbPage := total / displayed
	utils.Print(utils.Coordinates{X: 5, Y: h - 1}, defaultFooterAttributes, fmt.Sprintf("Pages: %d/%d", pageNb, nbPage))
	mangasCountFormat := "Mangas: %d"
	sizeOfMessage := float64(len(mangasCountFormat) + 5)
	marginRight := int(math.Min(math.Max(sizeOfMessage, float64(w/8)), float64(colSize)))
	utils.Print(utils.Coordinates{X: w - marginRight, Y: h - 1}, defaultFooterAttributes, fmt.Sprintf(mangasCountFormat, total))
}
