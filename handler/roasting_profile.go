package handler

import (
	"github.com/go-chi/chi/v5"
	"html/template"
	"instagram-roasting/core/module"
	"net/http"
)

type RoastingProfileHandler struct {
	roastingUC *module.RoastingUC
}

func NewRoastingProfileHandler(r chi.Router, scraperUC *module.RoastingUC) {
	handler := &RoastingProfileHandler{
		roastingUC: scraperUC,
	}
	r.Get("/roasting-ig/{username}", handler.GetRoastingProfile)
}

func (h *RoastingProfileHandler) GetRoastingProfile(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	result, err := h.roastingUC.GetRoastedProfile(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse and execute the template
	tmpl, err := template.ParseFiles("template/roasting.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Roast string
	}{
		Roast: result,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
