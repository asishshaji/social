package profile

type Mode string
type Driver string

const (
	DEV  Mode = "DEV"
	PROD Mode = "PROD"
)

const (
	POSTGRESQL Driver = "POSTGRESQL"
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
}

func GetProfile() (*Profile, error) {
	profile := new(Profile)

	profile.Addr = "localhost"
	profile.Port = 8080
	profile.Driver = POSTGRESQL

	return profile, nil
}
