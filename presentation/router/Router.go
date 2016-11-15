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
	departmentHandler := handlers.NewDepartmentHandler()

	router := httprouter.New()
	router.GET("/", handlers.IndexHandler)
	router.GET("/departments", departmentHandler.GetDepartments)
	router.POST("/departments", departmentHandler.CreateDepartment)
	/*router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/todos", handlers.TodoIndex)
	router.HandleFunc("/todos/{todoId}", handlers.TodoShow)*/

	return router
}
