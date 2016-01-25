package mangaeden

import (
  "encoding/json"
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"
)

type MangaAPI struct {
  I string `json: i`
  T string `json: t`
  C []string `json: c`
  Im string `json: im`
  A string `json: a`
  S int `json: s`
}

type mangaResponse struct {
  Manga []MangaAPI
}

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
