package main

import (
	"net/http"
	"Rein/config"
	log "github.com/Sirupsen/logrus"
	"fmt"
	"strconv"
)

type App struct {
	logger config.Logger
	port int
}

func main() {
	app := App{logger:config.Logger{}, port: 8080}
	router:= config.Router{}.NewRouter()

	app.logger.Execute()

	fmt.Printf("The server is running in port: " + strconv.Itoa(app.port) + "\n")
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(app.port), router))
}