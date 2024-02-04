package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go2api/pkg/models"
	"go2api/pkg/templates"
)

type Handler struct {
	Template templates.Templater
}

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome :)")
}

func (h *Handler) FindPlacesFirstPageHandler(w http.ResponseWriter, r *http.Request) {
	h.Template.ExecuteTemplate(w, "index.html", nil)
}

func (h *Handler) FindPlacesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	place := r.URL.Query().Get("place")
	area := r.URL.Query().Get("area")

	responseData := models.FindPlaces(place, area)

	json.NewEncoder(w).Encode(responseData)
}
