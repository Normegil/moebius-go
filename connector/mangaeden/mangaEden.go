// Package mangaeden implements the MangaEden API (http://www.mangaeden.com/api/).
package mangaeden

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// MangaAPI is the structure representing the response from MangaEden when loading the manga list.
// It represents only the list of mangas
type MangaAPI struct {
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

type mangaResponse struct {
	Manga []MangaAPI
}

// LoadMangas is loading the manga list from MangaEden.
// It accept a language parameter which should be "en"(English) or "it"(Italian). If the string is empty, it will default to english.
func LoadMangas(language string) ([]MangaAPI, error) {
	languageCode, err := getLanguageCode(language)
	if nil != err {
		return nil, err
	}

	const mangaURL = "https://www.mangaeden.com/api/list/%d/"
	url := fmt.Sprintf(mangaURL, languageCode)
	resp, err := http.Get(url)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	mangaResp, err := parseResponse(resp)
	if nil != err {
		return nil, err
	}
	return mangaResp.Manga, nil
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

func parseResponse(resp *http.Response) (mangaResponse, error) {
	content, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return mangaResponse{}, err
	}

	var mangaResp mangaResponse
	err = json.Unmarshal(content, &mangaResp)
	if nil != err {
		return mangaResponse{}, err
	}
	return mangaResp, nil
}
