package main

import (
	"net/http"

	"github.com/hdmorales87/go-gorm-restapi/db"
	_ "github.com/hdmorales87/go-gorm-restapi/docs"
	"github.com/hdmorales87/go-gorm-restapi/models"
	"github.com/hdmorales87/go-gorm-restapi/router"
)

// @title Go GORM REST API
// @version 1.0
// @description REST API construida con Go, GORM, Gorilla Mux y PostgreSQL para gestionar usuarios y tareas.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	database := db.Connect()
	database.AutoMigrate(&models.User{}, &models.Task{})

	mux := router.SetupRoutes(database)
	http.ListenAndServe(":8080", mux)
}
