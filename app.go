package main

import (
	"os"
	"os/user"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/normegil/moebius-go/cache"
	"github.com/normegil/moebius-go/connector"
	"github.com/normegil/moebius-go/connector/mangaeden"
	"github.com/normegil/moebius-go/views"
	"github.com/normegil/moebius-go/views/terminal/gui"
)

var logFile *os.File

func init() {
	logFile := initLogFile()
	log.SetOutput(logFile)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func main() {
	defer logFile.Close()
	err := gui.Launch(views.ViewInputs{
		Cache:   getCache(),
		Listers: getListers(),
	})
	if nil != err {
		log.Panic(err)
		panic(err)
	}
}

func getListers() []connector.Lister {
	var mangaEdenAPI mangaeden.API
	listers := []connector.Lister{mangaEdenAPI}
	return listers
}

func getCache() cache.Cache {
	c, err := cache.NewFileCache()
	if nil != err {
		log.Panic(err)
		panic(err)
	}
	c = &cache.FileCache{
		Path:       c.Path,
		Expiration: 5 * 24 * time.Hour,
	}
	return c
}

func getLogPath() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir + "/.moebius/log/moebius.log.json"
}

func initLogFile() *os.File {
	logFile, err := os.Create(getLogPath())
	if err != nil {
		panic(err)
	}
	return logFile
}
