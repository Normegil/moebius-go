package cache

import (
	"github.com/normegil/moebius-go/models"
)

// Saver save given Manga to a data store, depending on the implementation used
type Saver interface {
	Save([]models.Manga, string) error
}
