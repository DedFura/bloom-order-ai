package app

import (
	"bloom-order-ai/internal/config"
	"bloom-order-ai/internal/logging"
	"bloom-order-ai/internal/router"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Run(cfg *config.Config, log logging.Logger) {
	//dbConnector, err := data.NewDBConnector(log, cfg.DB.URL, cfg.DB.Log)
	//if err != nil {
	//	log.WithFields(logrus.Fields{
	//		"method": "NewDBConnector",
	//		"type":   "DatabaseError",
	//	}).WithError(err).Error("Failed to initialize database")
	//}

	r := router.NewRouter(cfg, log)

	//listingRepo := repository.NewRepository(dbConnector)

	c := setupCors()

	log.Infof("server started on %v", cfg.Addr)

	if err := http.ListenAndServe(cfg.Addr, c.Handler(r)); err != nil {
		log.WithFields(logrus.Fields{
			"type": logging.NetworkError,
		}).WithError(err).Error("Failed to start server")
	}
}

func setupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*", "ws://*", "wss://*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
}
