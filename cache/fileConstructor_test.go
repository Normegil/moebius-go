package cache

import (
	"os/user"
	"testing"
)

const errorMsg = "Test: %s\n\t%s"

func TestFileConstructPath(t *testing.T) {
	testName := "Test constructing FileCache for User home folder"
	cache, err := NewFileCache()
	if nil != err {
		t.Fatalf(errorMsg, testName, err)
	}

	usr, err := user.Current()
	if err != nil {
		t.Fatalf(errorMsg, testName, err)
	}
	path := usr.HomeDir + "/.moebius/cache"
	if path != cache.Path {
		errorMsgImproved := errorMsg + "\n\t\tExpected Path: %+v\n\t\tPath received: %+v"
		t.Fatalf(errorMsgImproved, testName, "Path not corresponding", path, cache.Path)
	}
}

func TestFileConstructExpiration(t *testing.T) {
	testName := "Test constructing FileCache for User home folder"
	cache, err := NewFileCache()
	if nil != err {
		t.Fatalf(errorMsg, testName, err)
	}

	if 0 != cache.Expiration {
		errorMsgImproved := errorMsg + "\n\t\tExpected Expiration: %s\n\t\tExpiration received: %s"
		t.Fatalf(errorMsgImproved, testName, "Path not corresponding", 0, cache.Expiration)
	}
}
