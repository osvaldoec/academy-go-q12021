package repository

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"

	"pokemon/entity"
)

type repo struct{}

// NewCsvPokemonRepository - this func helps to create a new repo
func NewCsvPokemonRepository() PokemonRepository {
	return &repo{}
}

func (*repo) FindOne(param string) ([]entity.Pokemon, error) {
	file, err := os.Open("test.csv")
	if err != nil {
		return nil, errors.New("Error: failed to open CSV file")
	}
	defer file.Close()
	var result []entity.Pokemon

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if record[0] == param {
			v1, err := strconv.Atoi(record[0])
			if err != nil {
				log.Println(err)
			}
			pokemon := entity.Pokemon{
				ID:             v1,
				Name:           record[1],
				BaseExperience: record[2],
			}
			result = append(result, pokemon)
			break
		}
	}
	if len(result) == 0 {
		return nil, errors.New("Pokemon not found")
	}
	return result, nil
}
