package main

import (
	"net/http"
	"truck-management/truck-management/infrastructure"
)

func main() {
	logger := infrastructure.LoggerFactory()

	config, cErr := infrastructure.NewConfig()
	if cErr != nil {
		logger.Fatal(cErr)
	}

	db, dErr := infrastructure.DatabaseFactory(config)
	if dErr != nil {
		logger.Fatal(dErr)
	}

	ts := infrastructure.TruckServiceFactory(db)
	h := infrastructure.HandlerFactory(ts, logger)

	logger.Info("starting webserver")

	if err := http.ListenAndServe(config.AppPort, h); err != nil {
		logger.Fatal(err)
	}
}
