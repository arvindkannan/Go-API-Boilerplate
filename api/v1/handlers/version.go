package handlers

import (
	"github.com/gin-gonic/gin"

	"go-api-boilerplate/types"
	"go-api-boilerplate/version"
)

// VersionHandler handles the version route
func VersionHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"version": version.Version,
	})
}

func VersionErrorSampleHandler(c *gin.Context) {
	c.Error(
		&types.ValidationError{
			Message: "Validation Sample Error Message",
			Fields:  []string{"version"},
		},
	)

	return
}
