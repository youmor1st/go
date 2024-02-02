package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (a *SuperheroApp) GetSuperheroesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a.Superheroes)
}

func (a *SuperheroApp) GetSuperheroHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	name := params["name"]

	for _, superhero := range a.Superheroes {
		if superhero.Name == name {
			json.NewEncoder(w).Encode(superhero)
			return
		}
	}

	http.NotFound(w, r)
}
