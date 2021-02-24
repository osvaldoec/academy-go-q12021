package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	})
	router.HandleFunc("/get", getPokemons).Methods("GET")
	log.Print("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
