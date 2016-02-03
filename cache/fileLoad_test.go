package cache

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
	"time"

	"github.com/normegil/moebius-go/models"
)

func TestFileLoadWithEmptyPath(t *testing.T) {
	cache := FileCache{}
	_, err := cache.Load("FakeFile")
	if nil == err {
		t.Fatalf("Test: %s\n\t%s", "Check load with Empty path", "Error was expected but none was received")
	}
}

func TestFileLoad(t *testing.T) {
	path, err := ioutil.TempDir(tempDirRoot, tempDirPrefix)
	if nil != err {
		t.Fatal(err)
	}
	cache := &FileCache{Path: path}

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

	const errorMsg = "%s\n\tPath: %s\n\tMangas: %+v"
	fileName, err := createTestFile(cache.Path, mangas)
	if nil != err {
		t.Fatalf(errorMsg, err, cache.Path, mangas)
	}
	filePath := cache.Path + "/" + fileName + ".json"

	loadedMangas, err := cache.Load(fileName)
	if nil != err {
		fatalf(filePath, t, errorMsg, err, filePath, mangas)
	}

	if !reflect.DeepEqual(loadedMangas, mangas) {
		fatalf(filePath, t, "%s\n\tExpected: %+v\n\tGot: %+v", "Loaded Mangas and expected mangas not corresponding", mangas, loadedMangas)
	}
}

func TestFileLoadInexistingData(t *testing.T) {
	path, err := ioutil.TempDir(tempDirRoot, tempDirPrefix)
	if nil != err {
		t.Fatal(err)
	}
	cache := &FileCache{Path: path}
	fileName := "NotExisting"
	filePath := path + "/" + fileName + ".json"
	_, err = cache.Load(fileName)
	derr, ok := err.(DataNotFoundError)
	if !ok {
		fatalf(filePath, t, "Error is not of requested type: %s", err)
	}
	if derr.ID != filePath {
		fatalf(filePath, t, "DataNotFoundError doesn't contains right key.\n\tExpected: %s\n\tGot: %s", fileName, derr.ID)
	}
}

func TestFileLoadIgnoreExpirationIfZeroValue(t *testing.T) {
	path, err := ioutil.TempDir(tempDirRoot, tempDirPrefix)
	if nil != err {
		t.Fatal(err)
	}
	cache := &FileCache{
		Path:                  path,
		Expiration:            0,
		lastModificationTimer: testLastModificationTimer{time.Time{}},
	}
	_, err = cache.Load("File")
	if nil != err {
		_, ok := err.(DataNotFoundError)
		if !ok {
			t.Fatalf("Only DataNotFoundError allowed: %s", err)
		}
	}
}

func TestFileLoadCheckExpirationIfSet(t *testing.T) {
	path, err := ioutil.TempDir(tempDirRoot, tempDirPrefix)
	if nil != err {
		t.Fatal(err)
	}
	cache := &FileCache{
		Path:                  path,
		Expiration:            1 * time.Hour,
		lastModificationTimer: testLastModificationTimer{time.Time{}},
	}
	_, err = cache.Load("File")

	if nil == err {
		t.Fatal("Error is nil")
	}

	_, ok := err.(ExpiredError)
	if !ok {
		t.Fatalf("Error is not of requested type: %s", err)
	}
}

func TestFileLoadCheckExpirationInFuture(t *testing.T) {
	path, err := ioutil.TempDir(tempDirRoot, tempDirPrefix)
	if nil != err {
		t.Fatal(err)
	}
	cache := &FileCache{
		Path:                  path,
		Expiration:            1 * time.Hour,
		lastModificationTimer: testLastModificationTimer{time.Now().Add(1 * time.Hour)},
	}
	_, err = cache.Load("File")

	if nil == err {
		t.Fatal("Error is nil")
	}

	_, ok := err.(ExpiredError)
	if !ok {
		t.Fatalf("Error is not of requested type: %s", err)
	}
}

func createTestFile(folderPath string, mangas []models.Manga) (string, error) {
	fileName := "TestLoad"
	filePath := folderPath + "/" + fileName + ".json"

	jsonDatas, err := json.Marshal(mangas)
	if nil != err {
		return "", err
	}

	const fileMode = 0644
	err = ioutil.WriteFile(filePath, []byte(jsonDatas), fileMode)
	return fileName, err
}

type testLastModificationTimer struct {
	ModificationDate time.Time
}

func (lastModifTimer testLastModificationTimer) LastModificationTime(path string) (time.Time, error) {
	return lastModifTimer.ModificationDate, nil
}
