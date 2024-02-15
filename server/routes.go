package server

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go-api-boilerplate/api/v1/routes"
	"go-api-boilerplate/internal"
)

func (s *Server) GenerateRoutes() http.Handler {
	var origins []string
	if o := os.Getenv("ENDPOINT_ORIGINS"); o != "" {
		origins = strings.Split(o, ",")
	}

	config := cors.DefaultConfig()
	config.AllowWildcard = true
	config.AllowBrowserExtensions = true

	// if env var is not set, config.AllowAllOrigin=true
	if len(origins) == 0 {
		config.AllowAllOrigins = true
	} else {
		config.AllowOrigins = origins
	}

	// Create a new router
	r := gin.Default()
	r.Use(
		cors.New(config),
		func(c *gin.Context) {
			c.Set("workDir", s.WorkDir)
			c.Next()
		},
	)
	r.Use(internal.WrapError)

	// Create a router group with the "/api" prefix
	apiGroup := r.Group("/api")

	// Define your routes here by loop through the routes directors and for each route file, add the route to the router
	routes := []func(*gin.RouterGroup){
		// v1, v2 routes
		routes.SetupVer1Route,
	}

	for _, route := range routes {
		route(apiGroup)
	}

	// return route to start the server
	return r

}
