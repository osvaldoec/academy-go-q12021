package main

import (
	"encoding/json"
	"net/http"
	"pokemon/repository"
)

var (
	repo repository.PokemonRepository = repository.NewPokemonRepository()
)

func getPokemons(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	keys, ok := req.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		resp.WriteHeader(http.StatusBadGateway)
		resp.Write([]byte(`{"error": "URL Param 'id'is missing"}`))
		return
	}
	key := keys[0]
	pokemons, err := repo.FindOne(key)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the pokemon"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(pokemons)
}
