package main

import (
	"github.com/joegasewicz/gomek"
	"joegasewicz/fantasy-banking-app/web/views"
	"net/http"
)

const (
	PORT = 5000
)

func init() {

}

func main() {
	//whitelist := [][]string{
	//	{
	//		"/", "GET",
	//	},
	//}

	// Gomek config
	c := gomek.Config{
		BaseTemplateName: "layout",
		BaseTemplates: []string{
			"./templates/layout.gohtml",
			"./templates/partials/footer.gohtml",
			"./templates/partials/navbar.gohtml",
			"./templates/partials/scripts.gohtml",
		},
	}

	// Create app & setup static files
	app := gomek.New(c)
	files := http.FileServer(http.Dir("static"))
	app.Handle("/static/", http.StripPrefix("/static/", files))

	// Views
	app.Route("/").View(views.GetHome).Methods("GET").Templates("./templates/views/home.gohtml")

	// Middleware
	app.Use(gomek.Logging)
	app.Use(gomek.CORS)
	//app.Use(gomek.Authorize(whitelist, func(r *http.Request) (bool, context.Context) {
	//	return true, nil
	//}))

	// App
	app.Listen(8000)
	app.Start()
}
