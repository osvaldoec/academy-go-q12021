package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"pokemon/entity"
	"pokemon/errors"

	"github.com/gorilla/mux"
)

// UseCase - interface to handle communication
type UseCase interface {
	GetByID(pokemonID string) (*entity.Pokemon, error)
	InsertByID(pokemonID string) (*entity.Pokemon, error)
	GetItermsPerWorker(numType string, items int, itemsPerWorkers int) (*[]entity.Pokemon, error)
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
	pokemonID := mux.Vars(req)["id"]

	pokemon, err := p.UseCase.GetByID(pokemonID)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err.Error()})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(pokemon)
}

// InsertByID - handle logic from requests and responses
func (p *Pokemon) InsertByID(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	pokemonID := mux.Vars(req)["pokemonId"]

	pokemon, err := p.UseCase.InsertByID(pokemonID)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err.Error()})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(pokemon)
}

// GetItermsPerWorker - get Items from the csv
func (p *Pokemon) GetItermsPerWorker(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	params := mux.Vars(req)
	fmt.Print(params["type"])
	if params["type"] != "even" && params["type"] != "odd" {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Invalid request"})
		return
	}

	numType := params["type"]
	numItems, _ := strconv.Atoi(params["items"])
	itemsPerWorker, _ := strconv.Atoi(params["items-per-worker"])

	pokemon, err := p.UseCase.GetItermsPerWorker(numType, numItems, itemsPerWorker)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err.Error()})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(pokemon)

}
