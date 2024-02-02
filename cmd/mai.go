package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/youmor1st/goland/internal/app"
)

func main() {
	r := mux.NewRouter()

	superheroApp := app.NewSuperheroApp()

	r.HandleFunc("/superheroes", superheroApp.GetSuperheroesHandler).Methods("GET")
	r.HandleFunc("/superheroes/{name}", superheroApp.GetSuperheroHandler).Methods("GET")

	fmt.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
