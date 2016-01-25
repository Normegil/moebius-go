package connector

import "github.com/normegil/moebius-go/connector/mangaeden"

// Manga represents manga general infos.
type Manga struct {
	ID     string
	Title  string
	Tags   []string
	Image  string // URL to main image representing the manga
	Alias  string
	Status int
}

// LoadMangas call all different APIs and return a list of mangas from different origins.
// Language can be "en" or "it". If it's empty, it will be replaced by "en".
func LoadMangas(language string) ([]Manga, error) {
	mangasFromMangaEden, err := mangaeden.LoadMangas(language)
	if nil != err {
		return nil, err
	}
	mangas := make([]Manga, 0, len(mangasFromMangaEden))
	for _, manga := range mangasFromMangaEden {
		mangas = append(mangas, Manga{
			ID:     manga.I,
			Title:  manga.T,
			Tags:   manga.C,
			Image:  manga.Im,
			Alias:  manga.A,
			Status: manga.S,
		})
	}
	return mangas, nil
}
