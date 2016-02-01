package main

import (
	"fmt"

	"github.com/normegil/moebius-go/connector"
	"github.com/normegil/moebius-go/connector/mangaeden"
)

func main() {
	var mangaEdenAPI mangaeden.API
	listers := []connector.Lister{mangaEdenAPI}
	mangas, err := connector.ListMangas(listers, "en")
	if nil != err {
		panic(err)
	}
	for _, manga := range mangas {
		fmt.Println(manga.Title)
	}
}
