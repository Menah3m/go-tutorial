package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"log"
	"net/http"
)

type hello struct {
	app.Compo
}

func (h *hello) Render() app.UI {
	return app.H1().Text("HEllo app .")
}

func main() {
	app.Route("/", &hello{})
	app.RunWhenOnBrowser()
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An hello world",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
