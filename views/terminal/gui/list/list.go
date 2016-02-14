package list

import (
	"fmt"
	"sort"

	"github.com/normegil/moebius-go/connector"
	"github.com/normegil/moebius-go/models"
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
	content  []models.Manga
	selected int
}

// Init initialize and load needed data into Lister
func (list *Lister) Init(args views.ViewInputs) error {
	utils.Print(utils.Coordinates{X: 50, Y: 0}, utils.Attributes{Foreground: termbox.ColorBlack, Background: termbox.ColorRed}, fmt.Sprintf("Test: %d", 1))
	mangas, err := connector.ListMangas(args.Cache, args.Listers, connector.ListMangasOptions{
		UseCache: true,
		Language: "en",
	})
	if nil != err {
		return err
	}
	list.content = make([]models.Manga, len(mangas))
	copy(list.content, mangas)
	sort.Sort(models.ByTitle(list.content))
	list.selected = 0
	return nil
}
