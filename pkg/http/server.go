package http

import (
	"fmt"
	"net/http"

	"go2api/pkg/routes"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

const (
	GET string = http.MethodGet
)

func initRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", routes.IndexHandler).
		Methods(GET)
	router.HandleFunc("/places", routes.FindPlacesHandler).
		Queries("place", "{place-type}", "area", "{area-name}").
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
