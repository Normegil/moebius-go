package list

import (
	"fmt"
	"math"

	"github.com/normegil/moebius-go/views/terminal/gui/utils"
	termbox "github.com/nsf/termbox-go"
)

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
	if total%displayed != 0 {
		nbPage++
	}
	utils.Print(utils.Coordinates{X: 5, Y: h - 1}, defaultFooterAttributes, fmt.Sprintf("Pages: %d/%d", pageNb, nbPage))
	mangasCountFormat := "Mangas: %d"
	sizeOfMessage := float64(len(mangasCountFormat) + 5)
	marginRight := int(math.Min(math.Max(sizeOfMessage, float64(w/8)), float64(colSize)))
	utils.Print(utils.Coordinates{X: w - marginRight, Y: h - 1}, defaultFooterAttributes, fmt.Sprintf(mangasCountFormat, total))
}
