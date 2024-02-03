package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go2api/pkg/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome :)")
}

func FindPlacesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	place := r.URL.Query().Get("place")
	area := r.URL.Query().Get("area")

	responseData := models.FindPlaces(place, area)

	json.NewEncoder(w).Encode(responseData)
}
