package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"pokemon/config"
	"pokemon/controller"
	"pokemon/router"
	"pokemon/service"
	"pokemon/usecase"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config-file", "config.yml", "Path")
	flag.Parse()

	configFileLoad, err := config.ReadConfig(configFile)
	if err != nil {
		log.Fatal("Error: ", err)
		os.Exit(1)
	}

	readFile, err := os.Open(configFileLoad.PokemonDB)
	if err != nil {
		log.Fatal("Failed to open pokemon csv file")
		os.Exit(1)
	}

	defer readFile.Close()

	// service
	pokemonService, err := service.New(readFile)

	// usecase
	usecase := usecase.New(pokemonService)

	// controller
	pokemonController := controller.New(usecase)

	// router
	router := router.NewRouter(pokemonController)

	fmt.Printf("HTTP server running on port %v", configFileLoad.Address)
	http.ListenAndServe(configFileLoad.Address, router)
}
