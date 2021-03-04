package controller

import (
	"encoding/json"
	"net/http"

	"pokemon/entity"
	"pokemon/errors"

	"github.com/gorilla/mux"
)

// UseCase - interface to handle communication
type UseCase interface {
	GetByID(pokemonID string) (*entity.Pokemon, error)
}

// Pokemon - struct to implement the usecase interface
type Pokemon struct {
	UseCase UseCase
}

// New - func to create a new pokemon controller
func New(u UseCase) *Pokemon {
	return &Pokemon{u}
}

// GetByID - handle logic from requests and responses
func (p *Pokemon) GetByID(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	pathParams := mux.Vars(req)

	pokemon, err := p.UseCase.GetByID(pathParams["id"])
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err.Error()})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(pokemon)
}
