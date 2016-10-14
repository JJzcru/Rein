package config

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

type Log struct{}

func (l Log) Execute() {
	l.init()
}

func (l Log) init() {
	switch os.Getenv("ENV") {
	case "production":
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.ErrorLevel)
	case "test":
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.DebugLevel)
	default:
		log.SetFormatter(&log.TextFormatter{ForceColors: true})
		log.SetLevel(log.InfoLevel)
	}
}

func (l Log) Fatal(args ...interface{}) {
	log.Fatal(args)
}
