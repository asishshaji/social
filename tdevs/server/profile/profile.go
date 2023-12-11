package profile

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

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

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	// TODO read from .env
	profile := new(Profile)

	profile.Addr = os.Getenv("APP_URL")
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		return nil, err
	}
	profile.Port = port
	profile.Driver = Driver(os.Getenv("DB_TYPE"))
	profile.JWT_Secret = os.Getenv("JWT_SECRET")
	profile.DSN = fmt.Sprintf("user=%s dbname=%s password=pass sslmode=disable", "postgres", "tdevs")

	return profile, nil
}
