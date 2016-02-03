package connector

import (
	cachePkg "github.com/normegil/moebius-go/cache"
	"github.com/normegil/moebius-go/models"
	"github.com/normegil/moebius-go/utils"
)

const listAllCacheID = "MoebiusListOfAllMangas"

// ListMangasOptions gather inputs for ListMangas method. Doesn't list dependancies
type ListMangasOptions struct {
	Language string
	useCache bool
}

// ListMangas call all different APIs and return a list of mangas from different origins.
// Language can be "en" or "it". If it's empty, it will be replaced by "en".
func ListMangas(cache cachePkg.Cache, listers []Lister, options ListMangasOptions) ([]models.Manga, error) {
	if options.useCache {
		return listUsingCache(cache, listers, options.Language)
	}
	return listUsingListers(listers, options.Language)
}

func listUsingCache(cache cachePkg.Cache, listers []Lister, language string) ([]models.Manga, error) {
	mangas, err := cache.Load(listAllCacheID + language)
	if nil != err {
		switch err.(type) {
		case cachePkg.DataNotFoundError, cachePkg.ExpiredError:
			mangas, err := listUsingListers(listers, language)
			if nil != err {
				return nil, err
			}
			cache.Save(mangas, listAllCacheID+language)
			return mangas, nil
		default:
			return nil, err
		}
	}
	return mangas, nil
}

func listUsingListers(listers []Lister, language string) ([]models.Manga, error) {
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
