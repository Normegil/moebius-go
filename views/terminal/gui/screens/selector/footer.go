package selector

import (
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
	"github.com/normegil/moebius-go/views/terminal/gui/components/list"
	"github.com/normegil/moebius-go/views/terminal/gui/utils"
	termbox "github.com/nsf/termbox-go"
)

const (
	formatMangasCount = "Mangas: %d"
	formatPages       = "Pages: %d/%d"
)

var footerAttr = utils.Attributes{
	Foreground: termbox.ColorBlack,
	Background: termbox.ColorCyan,
}

//Footer manage the display of the footer under the list of mangas
func (selector *Selector) Footer() error {
	w, h := termbox.Size()

	row := h - 1
	utils.Fill(utils.Coordinates{X: 0, Y: row}, utils.Size{Width: w, Height: 1}, footerAttr)

	printPages(*selector.list, row)
	printCounter(*selector.list, row, w)
	return nil
}

func printPages(listComponent list.List, row int) {
	displayed := listComponent.Displayed()

	current := listComponent.Page() + 1

	listSize := len(listComponent.Displayable)
	total := 1
	if displayed != 0 {
		total = listSize / displayed
	}
	if listSize > displayed && total%displayed != 0 {
		total++
	}
	log.WithFields(log.Fields{
		"total":              total,
		"displayed":          displayed,
		"filtered list size": listSize,
	}).Debug("Calculating display of total of pages")

	utils.Print(utils.Coordinates{X: 5, Y: row}, footerAttr, fmt.Sprintf(formatPages, current, total))
}

func printCounter(listComponent list.List, row, width int) {
	minSize := float64(len(formatMangasCount) + 5)
	maxSize := float64(listComponent.CellSize.Width)
	adapter := float64(width / 8)
	marginRight := int(math.Min(maxSize, math.Max(minSize, adapter)))
	utils.Print(utils.Coordinates{X: width - marginRight, Y: row}, footerAttr, fmt.Sprintf(formatMangasCount, len(listComponent.Displayable)))
}
