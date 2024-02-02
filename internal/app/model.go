package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Superhero struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
	Age   int    `json:"age"`
}

type SuperheroApp struct {
	Superheroes []Superhero
}

func NewSuperheroApp() *SuperheroApp {
	data, err := ioutil.ReadFile("data/superheroes.json")
	if err != nil {
		log.Fatal("Error reading superheroes.json:", err)
	}

	var superheroes []Superhero
	err = json.Unmarshal(data, &superheroes)
	if err != nil {
		log.Fatal("Error unmarshalling superheroes:", err)
	}

	return &SuperheroApp{Superheroes: superheroes}
}
