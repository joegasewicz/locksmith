package main

import (
	"context"
	"github.com/joegasewicz/gomek"
	"github.com/joegasewicz/locksmith/models"
	"github.com/joegasewicz/locksmith/utilities"
	"github.com/joegasewicz/locksmith/views"
	"net/http"
)

func main() {
	config := utilities.Config
	utilities.DB.AutoMigrate(
		&models.Role{},
		&models.User{},
	)

	// seed db
	s := utilities.Seeder{}
	s.CreateRoles()
	s.CreateUser()

	var whiteList = [][]string{
		{
			"/health", "GET",
		},
		{
			"/users", "GET",
		},
		{
			"/users", "POST",
		},
		{
			"/login", "POST",
		},
		{
			"/health", "GET",
		},
	}

	c := gomek.Config{}
	app := gomek.New(c)

	// views
	app.Route("/health").View(views.Health).Methods("GET")

	app.Route("/login").View(views.Login).Methods("POST")
	app.Route("/users").Resource(&views.Users{}).Methods("POST", "GET", "PUT", "DELETE")

	// middleware
	app.Use(gomek.CORS)
	app.Use(gomek.Authorize(whiteList, func(r *http.Request) (bool, context.Context) {
		return true, nil
	}))
	app.Use(gomek.Logging)

	app.Listen(config.Port)
	app.Start()
}
