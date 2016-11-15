package router

import "github.com/julienschmidt/httprouter"
import "Rein/presentation/handlers"

// Router of the application
type Router httprouter.Router

// NewRouter Empty constructor of the Router
func NewRouter() *Router {
	return &Router{}
}

// GetRouter Get the router object for implementation
func (r *Router) GetRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", handlers.IndexHandler)
	router.GET("/department", handlers.NewDepartmentHandler().GetDepartments)
	router.POST("/department", handlers.NewDepartmentHandler().CreateDepartment)
	/*router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/todos", handlers.TodoIndex)
	router.HandleFunc("/todos/{todoId}", handlers.TodoShow)*/
	return router
}
