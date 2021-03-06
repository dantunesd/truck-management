package infrastructure

type Config struct {
	AppName string `envconfig:"APP_NAME" default:"Truck-Management"`
	AppPort string `envconfig:"APP_PORT" default:":3000"`
	DBConn  string `envconfig:"DB_CONN" default:"root:root@tcp(db:3306)/truck_management"`
}
