package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jsnfwlr/o11y"
)

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code      int
	Err       error
	RequestID string
}

// Allows StatusError to satisfy the error interface.
func (statusErr StatusError) Error() string {
	return statusErr.Err.Error()
}

func (statusErr StatusError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Code      int    `json:"code"`
		Error     string `json:"error"`
		RequestID string `json:"request_id"`
	}{
		Code:      statusErr.Code,
		Error:     statusErr.Error(),
		RequestID: statusErr.RequestID,
	})
}

// Status returns our HTTP status code.
func (statusErr StatusError) Status() int {
	return statusErr.Code
}

// String returns our HTTP status code.
func (statusErr StatusError) String() string {
	return fmt.Sprintf("{ \"code\": %d, \"error\": \"%s\", \"request_id\": \"%s\"}", statusErr.Code, statusErr.Err.Error(), statusErr.RequestID)
}

func NewStatusError(ctx context.Context, code int, err error) StatusError {
	r := o11y.GetRequestID(ctx)

	return StatusError{
		RequestID: r,
		Code:      code,
		Err:       err,
	}
}
