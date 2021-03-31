package service

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"strconv"

	"pokemon/entity"

	"gopkg.in/resty.v1"
)

// Service - to create service for the csv file
type Service struct {
	csvFile  *os.File
	client   *resty.Client
	csvWrite *os.File
}

// New - func to create a new pokemon service
func New(csvFile *os.File, csvWrite *os.File) *Service {
	client := resty.New().SetHostURL("https://pokeapi.co/api/v2/")
	return &Service{csvFile, client, csvWrite}
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
			baseExp, err := strconv.Atoi(record[2])
			if err != nil {
				log.Println(err)
			}
			pokemon := entity.Pokemon{
				ID:             pokemonID,
				Name:           record[1],
				BaseExperience: baseExp,
			}

			return &pokemon, nil
		}
	}
	return nil, errors.New("Pokemon not found")
}

// InsertByID - function insert response in the csv
func (s *Service) InsertByID(param string) (*entity.Pokemon, error) {
	csvW := csv.NewWriter(s.csvWrite)
	response, err := s.client.R().
		SetPathParams(map[string]string{"pokemonId": param}).
		SetHeader("Accept", "application/json").
		Get("pokemon/{pokemonId}")
	result := &entity.Pokemon{}

	if err != nil {
		return nil, err
	}
	bodyResp := response.Body()
	err = json.Unmarshal(bodyResp, result)
	if err != nil {
		return nil, err
	}
	var data [][]string
	data = append(data, []string{strconv.Itoa(result.ID), result.Name, strconv.Itoa(result.BaseExperience)})
	csvW.WriteAll(data)

	if err := csvW.Error(); err != nil {
		return nil, err
	}
	csvW.Flush()
	return result, nil
}
