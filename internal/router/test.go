package router

import (
	"bloom-order-ai/internal/router/service"
	"github.com/gorilla/mux"
)

func TestRouter(root *mux.Router) {
	r := root.PathPrefix("/suggested").Subrouter()

	balanceAuth := r.PathPrefix("/").Subrouter()
	balanceAuth.HandleFunc("/", service.Test).
		Methods("GET")
}
