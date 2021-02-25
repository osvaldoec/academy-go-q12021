package repository

import "pokemon/entity"

// PokemonRepository - this interface contains the methods for the pokemon logic
type PokemonRepository interface {
	FindOne(param string) ([]entity.Pokemon, error)
}
