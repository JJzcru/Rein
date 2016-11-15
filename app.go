package main

import (
	"Rein/presentation/log"
	"Rein/presentation/router"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/cors"
)

// App Object that represents the app
type App struct {
	log    *log.Log
	port   string
	router *router.Router
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := App{log: log.NewLog(), port: port, router: router.NewRouter()}

	app.log.Execute()

	fmt.Printf("The server is running in port: " + app.port + "\n")
	app.log.Fatal(http.ListenAndServe(":"+app.port, cors.Default().Handler(app.router.GetRouter())))
}
