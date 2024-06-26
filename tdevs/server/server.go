package server

import (
	"context"
	"fmt"
	v1 "tdevs/api/v1"
	"tdevs/server/profile"
	"tdevs/store"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	e       *echo.Echo
	Profile *profile.Profile
	Store   *store.Store
}

func NewServer(ctx context.Context, profile *profile.Profile, store *store.Store) (*Server, error) {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	s := Server{
		e:       e,
		Store:   store,
		Profile: profile,
	}

	rootGroup := e.Group("")
	apiV1Service := v1.NewAPIV1Service(profile, store)
	apiV1Service.Register(rootGroup)

	return &s, nil
}

func (s *Server) Start(ctx context.Context) error {
	return s.e.Start(fmt.Sprintf("%s:%d", s.Profile.Addr, s.Profile.Port))
}

func (s *Server) Shutdown(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := s.e.Shutdown(ctx); err != nil {
		fmt.Printf("failed to shutdown server: %v", err)
	}
	if err := s.Store.Close(); err != nil {
		fmt.Printf("failed to close database: %v", err)
	}

}
