package service

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"

	"pokemon/entity"
)

// Service - to create service for the csv file
type Service struct {
	csvFile *os.File
}

// New - func to create a new pokemon service
func New(csvFile *os.File) (*Service, error) {
	// todo client deliverable 2
	return &Service{csvFile}, nil
}

// GetByID - func to read the csv file and search by ID a pokemon
func (s *Service) GetByID(param string) (*entity.Pokemon, error) {
	r := csv.NewReader(s.csvFile)
	defer s.csvFile.Seek(0, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if record[0] == param {
			pokemonID, err := strconv.Atoi(record[0])
			if err != nil {
				log.Println(err)
			}
			pokemon := entity.Pokemon{
				ID:             pokemonID,
				Name:           record[1],
				BaseExperience: record[2],
			}

			return &pokemon, nil
		}
	}
	return nil, errors.New("Pokemon not found")
}
