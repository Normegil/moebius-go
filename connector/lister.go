package connector

import (
	"github.com/normegil/moebius-go/models"
	"github.com/normegil/moebius-go/utils"
)

// Lister define a function to load Mangas from a source
type Lister interface {
	List(fetcher utils.Fetcher, language string) ([]models.Manga, error)
}
