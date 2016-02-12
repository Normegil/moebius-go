package raw

import (
	"fmt"

	"github.com/normegil/moebius-go/connector"
	"github.com/normegil/moebius-go/views"
)

// Launch will launch the application and print the raw outputs to the terminal
func Launch(args views.ViewInputs) error {
	mangas, err := connector.ListMangas(args.Cache, args.Listers, connector.ListMangasOptions{
		UseCache: true,
		Language: "en",
	})
	if nil != err {
		return err
	}

	for _, manga := range mangas {
		fmt.Println(manga.Title)
	}
	return nil
}
