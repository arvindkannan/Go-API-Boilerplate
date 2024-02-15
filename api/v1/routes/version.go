package routes

import (
	"go-api-boilerplate/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

// VersionRoute is a function that returns a version route
func VersionRoute(r *gin.RouterGroup) {
	r.GET("/version", handlers.VersionHandler)
	r.GET("/version/errorsample", handlers.VersionErrorSampleHandler)
}
