package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Controller - interface to handle req flow
type Controller interface {
	GetByID(resp http.ResponseWriter, req *http.Request)
}

// NewRouter - func to handle endpoints
func NewRouter(controller Controller) *mux.Router {
	var muxDispatcher = mux.NewRouter()
	muxDispatcher.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	})
	muxDispatcher.HandleFunc("/get/pokemon/{id}", controller.GetByID).Methods("GET")

	return muxDispatcher
}
