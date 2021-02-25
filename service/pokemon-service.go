package service

import (
	"errors"

	"pokemon/entity"
	"pokemon/repository"
)

var repo repository.PokemonRepository

// PokemonService - interface to manage pokemon methods
type PokemonService interface {
	Validate(keys []string, ok bool) error
	FindOne(param string) ([]entity.Pokemon, error)
}

type service struct{}

// NewPokemonService - fun to create a new pokemon service
func NewPokemonService(repository repository.PokemonRepository) PokemonService {
	repo = repository
	return &service{}
}

func (*service) Validate(keys []string, ok bool) error {
	if !ok || len(keys[0]) < 1 {
		err := errors.New("URL Param 'id'is missing")
		return err
	}
	return nil
}

func (*service) FindOne(param string) ([]entity.Pokemon, error) {
	return repo.FindOne(param)
}
