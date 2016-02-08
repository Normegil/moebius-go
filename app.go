package main

import (
	"time"

	"github.com/normegil/moebius-go/cache"
	"github.com/normegil/moebius-go/connector"
	"github.com/normegil/moebius-go/connector/mangaeden"
	"github.com/normegil/moebius-go/views"
	"github.com/normegil/moebius-go/views/terminal/gui"
)

func main() {
	gui.Launch(views.ViewInputs{
		Cache:   getCache(),
		Listers: getListers(),
	})
}

func getListers() []connector.Lister {
	var mangaEdenAPI mangaeden.API
	listers := []connector.Lister{mangaEdenAPI}
	return listers
}

func getCache() cache.Cache {
	c, err := cache.NewFileCache()
	if nil != err {
		panic(err)
	}
	c = &cache.FileCache{
		Path:       c.Path,
		Expiration: 5 * 24 * time.Hour,
	}
	return c
}
