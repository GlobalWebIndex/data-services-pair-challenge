package infrastructure

import (
	"net/http"
)

// AudienceRouter is the HTTP router for the audience-related endpoints.
type AudienceRouter struct {
	createHandler  http.Handler
	getByIDHandler http.Handler
}

// NewAudienceRouter creates a new AudienceRouter with the provided AudienceService and CreateAudienceHandler.
func NewAudienceRouter(
	createHandler http.Handler,
	getByIDHandler http.Handler,
) *AudienceRouter {
	return &AudienceRouter{
		createHandler:  createHandler,
		getByIDHandler: getByIDHandler,
	}
}

// SetupRoutes sets up the routes for the audience-related endpoints.
func (router *AudienceRouter) SetupRoutes() {
	http.Handle("/audience", router.createHandler)
	http.Handle("/audiences/", router.getByIDHandler)
	// Add more routes for other audience-related endpoints here.
}
