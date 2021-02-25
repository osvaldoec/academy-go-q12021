package main

import (
	"fmt"
	"net/http"

	"pokemon/config"
	"pokemon/controller"
	router "pokemon/http"
	"pokemon/repository"
	"pokemon/service"
)

var (
	pokemonRepository repository.PokemonRepository = repository.NewCsvPokemonRepository()
	pokemonService    service.PokemonService       = service.NewPokemonService(pokemonRepository)
	pokemonController controller.PokemonController = controller.NewPokemonController(pokemonService)
	httpRouter        router.Router                = router.NewMuxRouter()
)

func main() {
	config.ReadConfig()
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	})
	httpRouter.GET("/get", pokemonController.GetPokemons)
	httpRouter.SERVE(config.C.Server.Address)
}
