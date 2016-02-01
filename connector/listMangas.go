package connector

import (
	"github.com/normegil/moebius-go/models"
	"github.com/normegil/moebius-go/utils"
)

// ListMangas call all different APIs and return a list of mangas from different origins.
// Language can be "en" or "it". If it's empty, it will be replaced by "en".
func ListMangas(listers []Lister, language string) ([]models.Manga, error) {
	var fetcher utils.HTTPFetcher
	var mangas []models.Manga
	for _, lister := range listers {
		loadedMangas, err := lister.List(fetcher, language)
		mangas = append(mangas, loadedMangas...)
		if nil != err {
			return nil, err
		}
	}
	return mangas, nil
}
