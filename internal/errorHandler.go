package internal

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-api-boilerplate/types"
)

func handleError(c *gin.Context, err error) {
	commonError := &types.CommonError{
		// Set default values (adapt as needed)
		Code:    http.StatusInternalServerError,
		ErrorID: 0,
		Message: "Internal server error",
		//Details: err.Error(),
	}
	// Handle specific error types (optional)
	switch error := err.(type) {
	case *types.DatabaseError:
		commonError.Code = http.StatusInternalServerError
		commonError.ErrorID = 1
		commonError.Message = error.Error()
	case *types.NotFoundError:
		commonError.Code = http.StatusNotFound
		commonError.ErrorID = 2
		commonError.Message = "Not found"
	case *types.UnauthorizedError:
		commonError.Code = http.StatusUnauthorized
		commonError.ErrorID = 3
		commonError.Message = "Unauthorized"
	case *types.ValidationError:
		commonError.Code = http.StatusBadRequest
		commonError.ErrorID = 4
		commonError.Message = "Validation failed"
		commonError.Details = error.Error()
	case *types.InternalServerError:
		commonError.Code = http.StatusInternalServerError
		commonError.ErrorID = 5
		commonError.Message = "Internal server error"
	case *types.BadRequestError:
		commonError.Code = http.StatusBadRequest
		commonError.ErrorID = 6
		commonError.Message = "Bad request"
	case *types.ForbiddenError:
		commonError.Code = http.StatusForbidden
		commonError.Message = "Forbidden"
		commonError.ErrorID = 7
	case *types.ConflictError:
		commonError.Code = http.StatusConflict
		commonError.Message = "Conflict"
		commonError.ErrorID = 8
	default:
		// Log detailed error for debugging
		slog.Info(fmt.Sprintf("Unhandled error: %v", error))

	}

	c.AbortWithStatusJSON(commonError.Code, commonError)
}

func WrapError(c *gin.Context) {
	c.Next()

	// Check if there are any errors to handle
	if c.Errors.Last() != nil && c.Errors.Last().Err != nil {
		err := c.Errors.Last().Err

		if err != nil {
			handleError(c, err)
		}
	}
}
