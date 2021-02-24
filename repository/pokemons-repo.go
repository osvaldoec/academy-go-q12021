package repository

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"pokemon/entity"
	"strconv"
)

type PokemonRepository interface {
	FindOne(param string) ([]entity.Pokemon, error)
}

type repo struct{}

func NewPokemonRepository() PokemonRepository {
	return &repo{}
}

func (*repo) FindOne(param string) ([]entity.Pokemon, error) {
	file, err := os.Open("test.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var result []entity.Pokemon

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			if pe, ok := err.(*csv.ParseError); ok {
				fmt.Println("Bad col", pe.Column)
				fmt.Println("Bad line", pe.Line)
				fmt.Println("Error reported", pe.Err)
				if pe.Err == csv.ErrFieldCount {
					continue
				}
			}
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
		return nil, errors.New("No se encontró el pokemon")
	}
	return result, nil
}
