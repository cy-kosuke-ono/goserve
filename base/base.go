package base

import (
	"flag"
	"log"
	"os"
)

type Args struct {
	LogPath   *string
	Port      *string
	LogFileFd *os.File
}

func New() Args {
	args := Args{}
	args.parse()
	args.openLogFile()
	return args
}

func (a *Args) parse() {
	a.LogPath = flag.String("l", "/var/log/goserve/access.log", "Access log file path")
	a.Port = flag.String("p", "80", "Using port number")
	flag.Parse()
}

func (a *Args) openLogFile() {
	var err error
	a.LogFileFd, err = os.OpenFile(*a.LogPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *Args) CloseLogFile() {
	if err := a.LogFileFd.Close(); err != nil {
		log.Fatal(err)
	}
}
