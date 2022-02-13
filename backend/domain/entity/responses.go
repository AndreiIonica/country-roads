package entity

import "fmt"

// Standardized http response
type Response struct {
	Message string
}

var NotFoundError Response
var SuccesResponse Response

// Returns a new formatted error for duplicate entities in the DB
// Useful for standardized responses
func NewDuplicateEntityErrorResponse(field string) Response {
	return Response{Message: fmt.Sprintf("This value is already taken: %s", field)}
}
