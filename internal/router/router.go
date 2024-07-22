package router

import (
	"bloom-order-ai/internal/config"
	"bloom-order-ai/internal/logging"
	"bloom-order-ai/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(cfg *config.Config, log logging.Logger) *mux.Router {
	r := mux.NewRouter()

	root := r.PathPrefix("/api").Subrouter()
	root.Use(middleware.Recoverer,
		middleware.LoggerMiddleware(log),
	)

	TestRouter(root)

	return r
}
