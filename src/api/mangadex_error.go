package api

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type ErrorResponse struct {
	Result string           `json:"result"`
	Errors []MangadexErrors `json:"errors"`
}

type MangadexErrors struct {
	ID      uuid.UUID `json:"id"`
	Status  int       `json:"status"`
	Title   string    `json:"title"`
	Detail  string    `json:"detail"`
	Context string    `json:"context"`
}

func (err *ErrorResponse) GetResult() string {
	return err.Result
}

func (err *ErrorResponse) GetErrors() string {
	var errors strings.Builder
	for _, e := range err.Errors {
		errors.WriteString(fmt.Sprintf("Status: %d\n %s: %s\n", e.Status, e.Title, e.Detail))
	}
	return errors.String()
}
