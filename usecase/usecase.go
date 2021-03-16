package usecase

import (
	"pokemon/entity"
)

// Service - interface to help with communication
type Service interface {
	GetByID(PokemonID string) (*entity.Pokemon, error)
}

// UseCase - struct to help with app flow
type UseCase struct {
	service Service
}

// New - Func to create a new usecase
func New(service Service) *UseCase {
	return &UseCase{service}
}

// GetByID - method to return a pokemon by an ID
func (u *UseCase) GetByID(pockemonID string) (*entity.Pokemon, error) {
	resp, err := u.service.GetByID(pockemonID)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
