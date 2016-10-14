package config
import (
	"github.com/gorilla/mux"
	"Rein/presentation/handlers"
)
type Router mux.Router

func (r Router) NewRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/todos", handlers.TodoIndex)
	router.HandleFunc("/todos/{todoId}", handlers.TodoShow)
	return router
}
