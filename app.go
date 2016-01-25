package main

import (
	"fmt"
	"github.com/normegil/moebius-go/connector"
)

func main() {
	mangas, err := connector.LoadMangas("en")
	if nil != err {
		panic(err)
	}
	for _, manga := range mangas {
		fmt.Println(manga.Title)
	}
}
