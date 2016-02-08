package list

import "github.com/normegil/moebius-go/models"

func format(manga models.Manga) string {
	toPrint := manga.Title
	if len(toPrint) > printSize {
		toPrint = toPrint[0:overflowPrintSize] + "..."
	}
	return toPrint
}
