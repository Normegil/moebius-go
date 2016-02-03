package cache

import (
	"os/user"
	"testing"
)

func TestFileConstruct(t *testing.T) {
	const errorMsg = "Test: %s\n\t%s"
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
