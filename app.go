package main

import (
	"fmt"
	"time"

	"github.com/normegil/moebius-go/cache"
	"github.com/normegil/moebius-go/connector"
	"github.com/normegil/moebius-go/connector/mangaeden"
)

func main() {
	c, err := cache.NewFileCache()
	if nil != err {
		panic(err)
	}
	c = &cache.FileCache{
		Path:       c.Path,
		Expiration: 5 * 24 * time.Hour,
	}

	var mangaEdenAPI mangaeden.API
	listers := []connector.Lister{mangaEdenAPI}

	mangas, err := connector.ListMangas(c, listers, connector.ListMangasOptions{
		Language: "en",
	})
	if nil != err {
		panic(err)
	}

	for _, manga := range mangas {
		fmt.Println(manga.Title)
	}
}
