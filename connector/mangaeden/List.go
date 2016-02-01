// Package mangaeden implements the MangaEden API (http://www.mangaeden.com/api/).
package mangaeden

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/normegil/moebius-go/models"
	"github.com/normegil/moebius-go/utils"
)

type mangaResponse struct {
	Manga []mangaAPI
}

type mangaAPI struct {
	I  string   // ID
	T  string   // Title
	C  []string // Tags
	Im string   // Main image
	A  string   // Alias
	/* Status:
	   -0: Suspended
	   -1: Ongoing
	   -2: Completed
	*/
	S int
}

func (manga mangaAPI) format() models.Manga {
	return models.Manga{
		ID:     manga.I,
		Title:  manga.T,
		Tags:   manga.C,
		Image:  manga.Im,
		Alias:  manga.A,
		Status: manga.S,
	}
}

// List is loading the manga list from MangaEden.
// It accept a language parameter which should be "en"(English) or "it"(Italian). If the string is empty, it will default to english.
func (API) List(fetcher utils.Fetcher, language string) ([]models.Manga, error) {
	languageCode, err := getLanguageCode(language)
	if nil != err {
		return nil, err
	}

	const mangaURL = "https://www.mangaeden.com/api/list/%d/"
	url := fmt.Sprintf(mangaURL, languageCode)
	content, err := fetcher.Fetch(url)
	if nil != err {
		return nil, err
	}

	var mangaResp mangaResponse
	err = json.Unmarshal(content, &mangaResp)
	if nil != err {
		return nil, err
	}

	mangas := make([]models.Manga, len(mangaResp.Manga))
	for index, manga := range mangaResp.Manga {
		mangas[index] = manga.format()
	}
	return mangas, nil
}

func getLanguageCode(language string) (int, error) {
	if "" == language || "en" == language {
		return 0, nil
	}
	if "it" == language {
		return 1, nil
	}
	return -1, errors.New("Language not supported: " + language)
}
