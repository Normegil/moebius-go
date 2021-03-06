package cache

import (
	"fmt"

	"github.com/normegil/moebius-go/models"
)

// Loader attempt to load Manga from data store (depending on interface).
// If a instance cannot be found, it will send back a DataNotFoundError.
// If the cache has expired, an ExpiredError will be returned.
type Loader interface {
	Load(string) ([]models.Manga, error)
}

// DataNotFoundError represent a loading error where data could not be found
type DataNotFoundError struct {
	ID string
}

func (e DataNotFoundError) Error() string {
	return fmt.Sprintf("Data not found: %s", e.ID)
}

// ExpiredError is returned if the cache has expired
type ExpiredError struct {
	ID string
}

func (e ExpiredError) Error() string {
	return fmt.Sprintf("Cache has expired: %s", e.ID)
}
