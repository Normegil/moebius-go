package mangaeden

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/normegil/moebius-go/models"
)

const errorMsg = "Test: %s\n\t%s\n\t\tInput: %+v\n\t\tOutput: %+v"

var enMangas = []models.Manga{
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

var itMangas = []models.Manga{
	{
		ID:     "56709f67719a16349215d5c3",
		Title:  "Silver Gravekeeper",
		Tags:   []string{"Action", "Comedy", "Fantasy", "Shounen", "Supernatural"},
		Image:  "06/0652f9d4bc85cb4a06d964c81fd251a893bed2ceb89dc2cf9327eac7.jpg",
		Alias:  "silver-gravekeeper",
		Status: 1,
	}, {
		ID:     "560dd721719a162ca4b513e0",
		Title:  "Moon-led Journey Across Another World",
		Tags:   []string{"Action", "Adventure", "Comedy", "Fantasy", "Shounen"},
		Image:  "0b/0b44ddbf5c8bcb5eab53ecb59a0c81c509467f51ccf0e1c68694611b.jpg",
		Alias:  "moon-led-journey-across-another-world",
		Status: 1,
	},
}

var incomplete = []models.Manga{
	{
		ID:     "5372389645b9ef5a0b1d20d8",
		Title:  "Flower Dream",
		Tags:   []string{"Sci-fi"},
		Image:  "",
		Alias:  "flower-dream",
		Status: 1,
	},
}

var tests = []struct {
	name     string
	input    Input
	expected Output
}{
	{"Fetcher send an error", Input{sendError: true}, Output{expectError: true}},
	{"Manga doesn't have an ID", Input{missingID: true}, Output{expectError: true}},
	{"Manga doesn't have a title", Input{missingTitle: true}, Output{expectError: true}},
	{"Manga is incomplete", Input{incomplete: true}, Output{mangas: incomplete}},
	{"Load with no language", Input{language: ""}, Output{mangas: enMangas}},
	{"Load with english language", Input{language: "en"}, Output{mangas: enMangas}},
	{"Load with italian language", Input{language: "it"}, Output{mangas: itMangas}},
	{"Load with not supported language", Input{language: "fr"}, Output{expectError: true}},
}

func TestList(t *testing.T) {
	for _, test := range tests {
		var api API
		fetcher := FakeFetcher{
			sendError:    test.input.sendError,
			missingID:    test.input.missingID,
			missingTitle: test.input.missingTitle,
			incomplete:   test.input.incomplete,
		}
		mangas, err := api.List(fetcher, test.input.language)
		if test.expected.expectError && nil == err {
			t.Fatalf(errorMsg, test.name, "Error is nil when it shouldn't.", test.input, test.expected)
		} else if !test.expected.expectError && nil != err {
			t.Fatalf(errorMsg, test.name, err, test.input, test.expected)
		} else if !reflect.DeepEqual(test.expected.mangas, mangas) {
			message := fmt.Sprintf("Got %+v", mangas)
			t.Fatalf(errorMsg, test.name, message, test.input, test.expected)
		}
	}
}

type FakeFetcher struct {
	sendError    bool
	missingID    bool
	missingTitle bool
	incomplete   bool
}

func (fetcher FakeFetcher) Fetch(url string) ([]byte, error) {
	if fetcher.sendError {
		return nil, errors.New("TestError")
	}

	if fetcher.missingID {
		const missingIDJSON = `{manga: [{
          "a": "flower-dream",
          "c": [
            "Sci-fi"
          ],
          "h": 962,
          "im": "ad/ad8dbe2c909de99899f1015a360f75e3ced31023672d6ff0d2b7547c.jpg",
          "ld": 1416420134.0,
          "s": 1,
          "t": "Flower Dream"
        }]}`
		return []byte(missingIDJSON), nil
	}
	if fetcher.missingTitle {
		const missingTitleJSON = `{manga: [{
          "a": "flower-dream",
          "c": [
            "Sci-fi"
          ],
          "h": 962,
          "i": "5372389645b9ef5a0b1d20d8",
          "im": "ad/ad8dbe2c909de99899f1015a360f75e3ced31023672d6ff0d2b7547c.jpg",
          "ld": 1416420134.0,
          "s": 1
        }]}`
		return []byte(missingTitleJSON), nil
	}
	if fetcher.incomplete {
		const incomplete = `{"manga": [{
          "a": "flower-dream",
          "c": [
            "Sci-fi"
          ],
          "h": 962,
          "i": "5372389645b9ef5a0b1d20d8",
          "ld": 1416420134.0,
          "s": 1,
					"t": "Flower Dream"
        }]}`
		return []byte(incomplete), nil
	}

	if "https://www.mangaeden.com/api/list/0/" == url {
		const enJSON = `{"manga": [{
          "a": "flower-dream",
          "c": [
            "Sci-fi"
          ],
          "h": 962,
          "i": "5372389645b9ef5a0b1d20d8",
          "im": "ad/ad8dbe2c909de99899f1015a360f75e3ced31023672d6ff0d2b7547c.jpg",
          "ld": 1416420134.0,
          "s": 1,
          "t": "Flower Dream"
        }, {
          "a": "kanai-kun",
          "c": [
            "Drama",
            "Slice of Life"
          ],
          "h": 260,
          "i": "54430be945b9ef3a6d5818cc",
          "im": "05/05d5df58e440371496b217f94cc4894abeba1671bd9edf0e7cd774a1.jpg",
          "ld": 1420483442.0,
          "s": 2,
          "t": "Kanai-kun"
        }]}`
		return []byte(enJSON), nil
	}
	if "https://www.mangaeden.com/api/list/1/" == url {
		const itJSON = `{"manga": [{
          "a": "silver-gravekeeper",
          "c": [
            "Action",
            "Comedy",
            "Fantasy",
            "Shounen",
            "Supernatural"
          ],
          "h": 30732,
          "i": "56709f67719a16349215d5c3",
          "im": "06/0652f9d4bc85cb4a06d964c81fd251a893bed2ceb89dc2cf9327eac7.jpg",
          "ld": 1453444005.0,
          "s": 1,
          "t": "Silver Gravekeeper"
        }, {
          "a": "moon-led-journey-across-another-world",
          "c": [
            "Action",
            "Adventure",
            "Comedy",
            "Fantasy",
            "Shounen"
          ],
          "h": 2834,
          "i": "560dd721719a162ca4b513e0",
          "im": "0b/0b44ddbf5c8bcb5eab53ecb59a0c81c509467f51ccf0e1c68694611b.jpg",
          "ld": 1443779426.0,
          "s": 1,
          "t": "Moon-led Journey Across Another World"
        }]}`
		return []byte(itJSON), nil
	}

	return nil, fmt.Errorf("Reached end of function: Case not forseen or handled \n\tFetcher: %+v\n\tURL: %s", fetcher, url)
}

type Input struct {
	sendError    bool
	missingID    bool
	missingTitle bool
	incomplete   bool
	language     string
}

type Output struct {
	expectError bool
	mangas      []models.Manga
}
