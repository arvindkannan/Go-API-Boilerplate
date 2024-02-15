package routes

import (
	"github.com/gin-gonic/gin"
)

// BaseRoute is the base route for the API
func SetupVer1Route(r *gin.RouterGroup) {
	v1 := r.Group("/v1")

	// Version route
	VersionRoute(v1)

	// Add more routes here
}
