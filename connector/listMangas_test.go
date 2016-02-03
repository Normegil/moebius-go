package connector

import (
	"errors"
	"reflect"
	"testing"

	"github.com/normegil/moebius-go/cache"
	"github.com/normegil/moebius-go/models"
	"github.com/normegil/moebius-go/utils"
)

const errorMsg = "Test: %s\n\t%s\n\t\tInput: %+v\n\t\tOutput: %+v"

var testMangas = []models.Manga{
	{
		ID:    "1",
		Title: "Test1",
	}, {
		ID:    "2",
		Title: "Test2",
	}, {
		ID:    "3",
		Title: "Test3",
	}, {
		ID:    "4",
		Title: "Test4",
	},
}

var tests = []struct {
	name     string
	in       []Lister
	expected output
}{
	{
		name:     "Send an error",
		in:       []Lister{FakeLister{sendError: true}},
		expected: output{expectError: true},
	},
	{
		name: "Send an error as soon as a lister fail",
		in: []Lister{
			FakeLister{mangas: []models.Manga{testMangas[0]}},
			FakeLister{sendError: true},
		},
		expected: output{expectError: true},
	},
	{
		name: "fuse manga from listers",
		in: []Lister{
			FakeLister{mangas: []models.Manga{testMangas[0]}},
			FakeLister{mangas: []models.Manga{testMangas[1]}},
		},
		expected: output{
			mangas: []models.Manga{testMangas[0], testMangas[1]},
		},
	},
	{
		name: "fuse list of mangas from listers",
		in: []Lister{
			FakeLister{mangas: []models.Manga{testMangas[0], testMangas[1]}},
			FakeLister{mangas: []models.Manga{testMangas[2], testMangas[3]}},
		},
		expected: output{
			mangas: testMangas,
		},
	},
}

func TestListMangas(t *testing.T) {
	for _, test := range tests {
		mangas, err := ListMangas(cache.NewMemoryCache(), test.in, ListMangasOptions{
			Language: "en",
		})
		if test.expected.expectError && nil == err {
			t.Fatalf(errorMsg, test.name, "An error was expected but none was received.", test.in, test.expected)
		} else if !test.expected.expectError {
			if nil != err {
				t.Fatalf(errorMsg, test.name, err, test.in, test.expected)
			} else if !reflect.DeepEqual(test.expected.mangas, mangas) {
				t.Fatalf(errorMsg, test.name, mangas, test.in, test.expected)
			}
		}
	}
}

type output struct {
	expectError bool
	mangas      []models.Manga
}

type FakeLister struct {
	mangas    []models.Manga
	sendError bool
}

func (lister FakeLister) List(fetcher utils.Fetcher, language string) ([]models.Manga, error) {
	if lister.sendError {
		return nil, errors.New("FakeError")
	}
	return lister.mangas, nil
}

func equal(expected []models.Manga, toTest []models.Manga) bool {
	if len(expected) != len(toTest) {
		return false
	}
	for index := range expected {
		if expected[index].ID != toTest[index].ID {
			return false
		}
	}
	return true
}
