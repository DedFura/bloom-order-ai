package main

import (
	"bloom-order-ai/internal/app"
	"bloom-order-ai/internal/config"
	"bloom-order-ai/internal/logging"
	"flag"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logging.NewLogrusLogger()

	defer func() {
		if r := recover(); r != nil {
			log.WithField("type", "PanicRecovery").Errorf("Recovered from panic: %v", r)
		}
	}()

	cfgPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.NewConfigFromFile(*cfgPath)
	if err != nil {
		log.WithFields(logrus.Fields{
			"type":   logging.ConfigurationError,
			"config": *cfgPath,
		}).WithError(err).Error("Failed to parse config file")
	}

	if err := log.SetLevel(cfg.Log); err != nil {
		log.WithFields(logrus.Fields{
			"type":   "ConfigurationError",
			"config": *cfgPath,
		}).WithError(err).Error("Failed to set log level")
	}

	app.Run(cfg, log)
}
