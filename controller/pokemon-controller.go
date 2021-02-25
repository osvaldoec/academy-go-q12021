package controller

import (
	"encoding/json"
	"net/http"

	"pokemon/errors"
	"pokemon/service"
)

var pokemonService service.PokemonService

// PokemonController - interface to implement differents pokemon methods
type PokemonController interface {
	GetPokemons(resp http.ResponseWriter, req *http.Request)
}
type controller struct{}

// NewPokemonController - func to create a new pokemon controller
func NewPokemonController(service service.PokemonService) PokemonController {
	pokemonService = service
	return &controller{}
}

func (*controller) GetPokemons(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	keys, ok := req.URL.Query()["id"]

	err := pokemonService.Validate(keys, ok)
	if err != nil {
		resp.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	key := keys[0]
	pokemons, err := pokemonService.FindOne(key)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err.Error()})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(pokemons)
}
