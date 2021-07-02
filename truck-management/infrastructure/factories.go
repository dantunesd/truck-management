package infrastructure

import (
	"net/http"
	"truck-management/truck-management/api"
	"truck-management/truck-management/application"

	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DatabaseFactory(config *Config) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(config.DBConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

func LoggerFactory() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&ecslogrus.Formatter{})
	return logger
}

func TruckServiceFactory(db *gorm.DB) *application.TruckService {
	return application.NewTruckService(
		TruckRepositoryFactory(db),
	)
}

func TruckRepositoryFactory(db *gorm.DB) application.ITruckRepository {
	return NewTruckRepository(db)
}

func HandlerFactory(db *gorm.DB, logger *logrus.Logger) http.Handler {
	return api.NewRouter(
		api.NewTruckHandler(TruckServiceFactory(db)),
		logger,
	).GetRoutes()
}
