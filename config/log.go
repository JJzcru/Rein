package config

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

type Logger struct{ }

func (l Logger) Execute(){
	l.init()
}

func (l Logger) init(){
	switch os.Getenv("ENV"){
	case "production":
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.ErrorLevel)
	case "test":
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.DebugLevel)
	default:
		log.SetFormatter(&log.TextFormatter{ForceColors:true})
		log.SetLevel(log.InfoLevel)
	}
}


