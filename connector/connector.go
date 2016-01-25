package connector

import "github.com/normegil/moebius-go/connector/mangaeden"

type Manga struct {
  ID string
  Title string
  Tags []string
  Image string
  Alias string
  Status int
}

func LoadMangas(language string) ([]Manga, error) {
  mangasFromMangaEden, err := mangaeden.LoadMangas(language)
  if nil != err {
    return nil, err
  }
  mangas := make([]Manga, 0, len(mangasFromMangaEden))
  for _, manga := range mangasFromMangaEden {
    mangas = append(mangas, Manga{
      ID: manga.I,
      Title: manga.T,
      Tags: manga.C,
      Image: manga.Im,
      Alias: manga.A,
      Status: manga.S,
    })
  }
  return mangas, nil
}
