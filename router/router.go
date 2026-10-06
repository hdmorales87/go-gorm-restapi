package router

import (
	"github.com/gorilla/mux"
	"github.com/hdmorales87/go-gorm-restapi/routes"
	"github.com/hdmorales87/go-gorm-restapi/services"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

func SetupRoutes(database *gorm.DB) *mux.Router {
	userService := services.NewUserService(database)
	taskService := services.NewTaskService(database)

	mux := mux.NewRouter()
	mux.HandleFunc("/", routes.HomeHandler())
	mux.HandleFunc("/users", routes.GetUsersHandler(userService)).Methods("GET")
	mux.HandleFunc("/users/{id}", routes.GetUserHandler(userService)).Methods("GET")
	mux.HandleFunc("/users", routes.CreateUserHandler(userService)).Methods("POST")
	mux.HandleFunc("/users/{id}", routes.UpdateUserHandler(userService)).Methods("PUT")
	mux.HandleFunc("/users/{id}", routes.DeleteUserHandler(userService)).Methods("DELETE")

	mux.HandleFunc("/tasks", routes.GetTasksHandler(taskService)).Methods("GET")
	mux.HandleFunc("/tasks/{id}", routes.GetTaskHandler(taskService)).Methods("GET")
	mux.HandleFunc("/tasks", routes.CreateTaskHandler(taskService)).Methods("POST")
	mux.HandleFunc("/tasks/{id}", routes.UpdateTaskHandler(taskService)).Methods("PUT")
	mux.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler(taskService)).Methods("DELETE")

	// Swagger documentation route
	mux.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return mux
}
