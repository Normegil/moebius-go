package cache

import (
	"reflect"
	"testing"

	"github.com/normegil/moebius-go/models"
)

func TestMemoryLoad(t *testing.T) {
	id := "Identifier"
	cache := NewMemoryCache()
	mangas := []models.Manga{
		{
			ID:     "5372389645b9ef5a0b1d20d8",
			Title:  "Flower Dream",
			Tags:   []string{"Sci-fi"},
			Image:  "ad/ad8dbe2c909de99899f1015a360f75e3ced31023672d6ff0d2b7547c.jpg",
			Alias:  "flower-dream",
			Status: 1,
		}, {
			ID:     "54430be945b9ef3a6d5818cc",
			Title:  "Kanai-kun",
			Tags:   []string{"Drama", "Slice of Life"},
			Image:  "05/05d5df58e440371496b217f94cc4894abeba1671bd9edf0e7cd774a1.jpg",
			Alias:  "kanai-kun",
			Status: 2,
		},
	}
	cache.memory[id] = mangas

	loadedMangas, err := cache.Load(id)
	if nil != err {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(mangas, loadedMangas) {
		t.Fatalf("%s\n\tExpected: %+v\n\tGot     : %+v", "Loaded mangas and mangas are differents", mangas, loadedMangas)
	}
}

func TestMemoryLoadInexistingData(t *testing.T) {
	id := "NotExiting"
	cache := NewMemoryCache()
	_, err := cache.Load(id)
	derr, ok := err.(DataNotFoundError)
	if !ok {
		t.Fatalf("Error is not of requested type: %s", err)
	}
	if derr.ID != id {
		t.Fatalf("DataNotFoundError doesn't contains right key.\n\tExpected: %s\n\tGot: %s", id, derr.ID)
	}
}
