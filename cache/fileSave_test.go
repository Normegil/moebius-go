package cache

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/normegil/moebius-go/models"
)

const tempDirRoot = ""
const tempDirPrefix = "githubcom_normegil_moebiusgo_cache_testsave"

type TestData struct {
	name        string
	emptyPath   bool
	expectError bool
	fileName    string
	mangas      []models.Manga
}

func TestFileSaveWithEmptyPath(t *testing.T) {
	cache := FileCache{}
	err := cache.Save([]models.Manga{}, "")
	if nil == err {
		t.Fatalf("Test: %s\n\t%s", "Check for empty path", "An error was expected but none was received.")
	}
}

func TestFileSave(t *testing.T) {
	const errorMsg = "Test: %s\n\t%s"
	testName := "Check that it save as expected"

	path, err := ioutil.TempDir(tempDirRoot, tempDirPrefix)
	if nil != err {
		t.Fatalf(errorMsg, testName, err)
	}
	cache := &FileCache{Path: path}

	fileName := "TestSave"
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

	err = cache.Save(mangas, fileName)
	filePath := cache.Path + "/" + fileName + ".json"
	errorMsgImproved := errorMsg + "\n\t\tFilePath: %+v\n\t\tMangas: %+v"
	if nil != err {
		fatalf(filePath, t, errorMsgImproved, testName, err, fileName, mangas)
	}

	jsonDatas, err := ioutil.ReadFile(filePath)
	if nil != err {
		fatalf(filePath, t, errorMsgImproved, testName, err, fileName, mangas)
	}

	var mangasLoaded []models.Manga
	err = json.Unmarshal(jsonDatas, &mangasLoaded)
	if nil != err {
		fatalf(filePath, t, errorMsgImproved, testName, err, fileName, mangas)
	}

	if !reflect.DeepEqual(mangas, mangasLoaded) {
		fatalf(filePath, t, errorMsgImproved, testName, string(jsonDatas), fileName, mangas)
	}
}

func fatalf(path string, t *testing.T, format string, args ...interface{}) {
	err := os.Remove(path)
	if nil != err {
		t.Fatalf("ERROR WHILE REMOVING: %s", path)
	}
	t.Fatalf(format, args...)
}
