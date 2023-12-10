package profile

import "fmt"

type Mode string
type Driver string

const (
	DEV  Mode = "DEV"
	PROD Mode = "PROD"
)

const (
	POSTGRESQL Driver = "postgres"
)

type Profile struct {
	// "prod", "dev"
	Mode Mode `json:"mode"`
	// server address
	Addr string `json:"addr"`
	// server port
	Port int `json:"port"`
	// database connection string
	DSN string `json:"dsn"`
	// database driver
	Driver Driver `json:"driver"`
	//application version
	Version string `json:"version"`
	//JWT secret
	JWT_Secret string `json:"secret"`
}

func GetProfile() (*Profile, error) {
	// TODO read from .env
	profile := new(Profile)

	profile.Addr = "localhost"
	profile.Port = 8080
	profile.Driver = POSTGRESQL
	profile.JWT_Secret = "secret"
	profile.DSN = fmt.Sprintf("user=%s dbname=%s password=pass sslmode=disable", "postgres", "tdevs")

	return profile, nil
}
