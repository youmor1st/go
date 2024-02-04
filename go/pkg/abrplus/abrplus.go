package abrplus

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetFriends(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, http.StatusOK, FriendsList)
}

func GetFriend(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	friend := GetFriendByName(name)
	if friend == nil {
		RespondWithError(w, http.StatusNotFound, "Friend not found")
		return
	}

	RespondWithJSON(w, http.StatusOK, friend)
}

func GetFriendByName(name string) *Friend {
	for _, friend := range FriendsList {
		if friend.Name == name {
			return &friend
		}
	}
	return nil
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	RespondWithJSON(w, statusCode, map[string]string{"error": message})
}
