package infrastructure

import (
	"truck-management/truck-management/application"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func LoggerFactory() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&ecslogrus.Formatter{})
	return logger
}

func ConfigFactory() (*Config, error) {
	cfg := &Config{}
	return cfg, envconfig.Process("", cfg)
}

func DatabaseFactory(config *Config) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(config.DBConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

func TruckServiceFactory(db *gorm.DB) *application.TruckService {
	return application.NewTruckService(NewTruckRepository(db))
}

func LocationServiceFactory(db *gorm.DB) *application.LocationService {
	return application.NewLocationService(NewLocationRepository(db))
}
