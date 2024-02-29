package main

import (
	"go2api/cmd/api"
	"go2api/pkg/db"
	"go2api/pkg/http"
)

func main() {
	api.RetrieveAPIKeys()
	db.PromptGeocodeDB()
	http.StartServer()
}
