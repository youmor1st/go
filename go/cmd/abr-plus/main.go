package abr_plus

import (
	"fmt"
	"github.com/codev0/go-dev-tsis1/pkg/abrplus"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/friends", abrplus.GetFriends).Methods("GET")
	r.HandleFunc("/friends/{name}", abrplus.GetFriend).Methods("GET")
	r.HandleFunc("/health", HealthCheck).Methods("GET")

	port := 8080
	fmt.Printf("Server is running on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Health check passed"))
}
