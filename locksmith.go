package main

import (
	"context"
	"github.com/joegasewicz/gomek"
	"github.com/joegasewicz/locksmith/models"
	"github.com/joegasewicz/locksmith/utilities"
	"github.com/joegasewicz/locksmith/views"
	"log"
	"net/http"
)

func main() {
	y := utilities.NewYaml()
	y.Get("locksmith.yaml")
	err := y.Do()
	if err != nil {
		log.Fatalln("error parsing locksmith.yaml")
	}
	config := utilities.Config
	utilities.DB.AutoMigrate(
		&models.Role{},
		&models.User{},
	)

	// seed db
	s := utilities.Seeder{}
	s.CreateRoles(&y.Yaml)
	s.CreateUser(&y.Yaml)

	var whiteList = [][]string{
		{
			"/", "GET",
		},
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
	}

	c := gomek.Config{
		BaseTemplateName: "layout",
		BaseTemplates: []string{
			"./templates/layout.gohtml",
			"./templates/partials/footer.gohtml",
			"./templates/partials/header.gohtml",
			"./templates/partials/navbar.gohtml",
			"./templates/partials/scripts.gohtml",
		},
	}
	app := gomek.New(c)

	// static files
	distFiles := http.FileServer(http.Dir("dist"))
	publicFiles := http.FileServer(http.Dir("public"))
	app.Handle("/dist/", http.StripPrefix("/dist/", distFiles))
	app.Handle("/public/", http.StripPrefix("/public/", publicFiles))

	// api views
	app.Route("/health").View(views.Health).Methods("GET")
	app.Route("/login").View(views.Login).Methods("POST")
	app.Route("/users").Resource(&views.Users{}).Methods("POST", "GET", "PUT", "DELETE")

	// template views
	app.Route("/").Resource(&views.Home{}).Methods("GET").Templates(
		"./templates/views/home.gohtml",
	)

	// middleware
	app.Use(gomek.CORS)
	app.Use(gomek.Authorize(whiteList, func(r *http.Request) (bool, context.Context) {
		return true, nil
	}))
	app.Use(gomek.Logging)

	app.Listen(config.Port)
	app.Start()
}
