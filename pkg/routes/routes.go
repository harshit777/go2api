package routes

import (
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
	w.Header().Set("Content-Type", "text/html")

	place := r.URL.Query().Get("place")
	area := r.URL.Query().Get("area")

	responseData := models.FindPlaces(place, area)

	err := h.Template.ExecuteTemplate(w, "places_list.html", responseData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
