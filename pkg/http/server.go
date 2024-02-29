package http

import (
	"fmt"
	"log"
	"net/http"

	"go2api/pkg/routes"
	"go2api/pkg/templates"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

const (
	GET string = http.MethodGet
)

func initRouter() *mux.Router {
	router := mux.NewRouter()

	tmpl := &templates.Template{}
	pattern := "static/*.html"
	err := tmpl.LoadTemplates(pattern)
	if err != nil {
		log.Fatal(err)
	}

	handler := &routes.Handler{
		Template: tmpl,
	}

	router.HandleFunc("/", handler.IndexHandler).
		Methods(GET)
	router.HandleFunc("/places", handler.FindPlacesHandler).
		Queries(
			"place", "{place}",
			"area", "{area}",
			"radius", "{radius:[0-9]+}",
			"sort", "{sort}",
		).
		Methods(GET)
	router.HandleFunc("/places", handler.FindPlacesFirstPageHandler).
		Methods(GET)

	return router
}

func StartServer() {
	mainRouter := initRouter()
	http.Handle("/", mainRouter)

	fmt.Println("Someone has entered your website")

	err := http.ListenAndServe(":8081", context.ClearHandler(http.DefaultServeMux))
	if err != nil {
		panic(err)
	}
}
