package types

// Define error types extending default error type
type NotFoundError struct {
	error
}

type UnauthorizedError struct {
	error
}

type ValidationError struct {
	error
	Message string   `json:"message"`
	Fields  []string `json:"fields,omitempty"`
}

type InternalServerError struct {
	error
}

type BadRequestError struct {
	error
}

type ForbiddenError struct {
	error
}

type ConflictError struct {
	error
}

type DatabaseError struct {
	error
	Message string `json:"message"`
}

type CommonError struct {
	Code    int    `json:"code"`
	ErrorID int    `json:"errorID,omitempty"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
	// Add more fields as needed (e.g., fields, more details etc.)
}

func (e *ValidationError) Error() string {
	return e.Message
}

func (e *DatabaseError) Error() string {
	return e.Message
}
