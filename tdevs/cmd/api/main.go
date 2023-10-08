package main

import (
	"log"
	"os"
	"tdevs/internal/cache"
	"tdevs/internal/user"
	"tdevs/util"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env file")
	}
	dbConn, err := util.ConnectToDB(util.DBOpts{
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USER:     os.Getenv("DB_USER"),
		SSLMODE:     os.Getenv("SSL_MODE"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_TYPE:     os.Getenv("DB_TYPE"),
	})
	if err != nil {
		log.Fatalf("db failed %s", err)
	}
	cache := cache.NewRedisCache(&redis.Options{})

	repo := user.NewUserRepo(log.Default(), dbConn)

	uS := user.NewUserService(repo, cache)
	uH := user.NewUserHandler(uS)

	e := echo.New()
	e.Use(middleware.Logger())

	api := e.Group("/api")

	//auth
	api.POST("/register", uH.RegisterUser)
	api.POST("/login", uH.Login)

	// restricted := api.Group("")
	// restricted.Use(echojwt.JWT([]byte("ASD")))

	e.Logger.Fatal(e.Start(":3000"))
}
