package cache

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/user"
	"time"

	"github.com/normegil/moebius-go/models"
	"github.com/normegil/moebius-go/utils"
)

// FileCache is a cache manager that will manage cache throught different files in the file system
type FileCache struct {
	Path string
	// Expiration is a duration after which the cache expire. If Expiration is 0, the cache will not expire.
	Expiration            time.Duration
	lastModificationTimer utils.LastModificationTimer
}

// NewFileCache creates a FileCache pointing to ".cache/" folder inside current user's home folder
// This cache doesn't expire
func NewFileCache() (*FileCache, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	path := usr.HomeDir + "/.moebius/cache"
	return &FileCache{Path: path}, nil
}

// Save will save a list of mangas in a file inside cache directory.
func (cache *FileCache) Save(mangas []models.Manga, fileName string) error {
	err := cache.checkSettings()
	if nil != err {
		return err
	}

	filePath := cache.Path + "/" + fileName + ".json"
	jsonDatas, err := json.Marshal(mangas)
	if nil != err {
		return err
	}

	const fileMode = 0644
	err = ioutil.WriteFile(filePath, []byte(jsonDatas), fileMode)
	return err
}

// Load will load a list of mangas from a file inside cache directory
func (cache *FileCache) Load(fileName string) ([]models.Manga, error) {
	err := cache.checkSettings()
	if nil != err {
		return nil, err
	}

	filePath := cache.Path + "/" + fileName + ".json"
	if 0 != cache.Expiration {
		err = manageExpiration(cache, filePath)
		if nil != err {
			return nil, err
		}
	}

	jsonData, err := ioutil.ReadFile(filePath)
	if nil != err {
		if os.IsNotExist(err) {
			return nil, DataNotFoundError{filePath}
		}
		return nil, err
	}

	var mangas []models.Manga
	err = json.Unmarshal(jsonData, &mangas)
	return mangas, err
}

func (cache *FileCache) checkSettings() error {
	if "" == cache.Path {
		return errors.New("Path was not initialized")
	}

	return nil
}

func manageExpiration(cache *FileCache, path string) error {
	lastModifTime, err := getLastModificationTime(cache, path)
	if nil != err {
		if os.IsNotExist(err) {
			return DataNotFoundError{path}
		}
		return err
	}

	now := time.Now()
	validityTime := now.Add(-cache.Expiration)
	if lastModifTime.Before(validityTime) || lastModifTime.After(now) {
		return ExpiredError{path}
	}
	return nil
}

func getLastModificationTime(cache *FileCache, path string) (time.Time, error) {
	if nil != cache.lastModificationTimer {
		return cache.lastModificationTimer.LastModificationTime(path)
	}
	stats, err := os.Stat(path)
	if nil != err {
		return time.Time{}, err
	}
	return stats.ModTime(), nil
}
