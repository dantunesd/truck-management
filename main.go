package main

import (
	"net/http"
	"truck-management/truck-management/api"
	"truck-management/truck-management/infrastructure"
)

func main() {
	logger := infrastructure.LoggerFactory()

	config, err := infrastructure.ConfigFactory()
	if err != nil {
		logger.Fatal(err)
	}

	db, err := infrastructure.DatabaseFactory(config)
	if err != nil {
		logger.Fatal(err)
	}

	handler := api.CreateHandler(
		logger,
		infrastructure.TruckServiceFactory(db),
		infrastructure.LocationServiceFactory(db),
		infrastructure.TripServiceFactory(db),
	)

	logger.Info("starting webserver")

	if err := http.ListenAndServe(config.AppPort, handler); err != nil {
		logger.Fatal(err)
	}
}
