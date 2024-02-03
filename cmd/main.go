package main

import (
	"go2api/cmd/api"
	"go2api/pkg/http"
)

func main() {
	api.RetrieveAPIKeys()
	http.StartServer()
}
