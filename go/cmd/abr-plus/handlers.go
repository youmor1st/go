package abr_plus

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetFriends(w http.ResponseWriter, r *http.Request) {
	friends := []string{"Friend1", "Friend2", "Friend3"}
	respondWithJSON(w, http.StatusOK, friends)
}

func GetFriend(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	friend := map[string]string{"name": name, "status": "Best friend"}
	respondWithJSON(w, http.StatusOK, friend)
}

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
