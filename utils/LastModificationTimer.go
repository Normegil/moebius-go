package utils

import (
	"time"
)

// LastModificationTimer fetch statistics for a given file
type LastModificationTimer interface {
	LastModificationTime(path string) (time.Time, error)
}
