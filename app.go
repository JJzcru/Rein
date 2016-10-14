package main

import (
	"Rein/config"
	"fmt"
	"net/http"
	"strconv"
)

type App struct {
	log  config.Log
	port int
}

func main() {
	app := App{log: config.Log{}, port: 8080}
	router := config.Router{}.NewRouter()

	app.log.Execute()

	fmt.Printf("The server is running in port: " + strconv.Itoa(app.port) + "\n")
	app.log.Fatal(http.ListenAndServe(":"+strconv.Itoa(app.port), router))
}
