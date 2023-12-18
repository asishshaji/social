package v1

import (
	"tdevs/server/profile"
	"tdevs/store"

	"github.com/labstack/echo/v4"
)

type APIV1Service struct {
	Profile *profile.Profile
	Store   *store.Store
}

func NewAPIV1Service(profile *profile.Profile, store *store.Store) *APIV1Service {
	return &APIV1Service{
		Profile: profile,
		Store:   store,
	}
}

func (s *APIV1Service) Register(rootGroup *echo.Group) {
	apiV1Group := rootGroup.Group("/api/v1")
	s.registerAuthRoutes(apiV1Group)
	s.registerUserRoutes(apiV1Group)
}
