package spotify

import (
	"fmt"
	"net/http"
)

var NoMorePagesError = &Error{
	Message: "spotify: no more pages",
	Status:  http.StatusBadRequest,
}

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Status code: %d. Message: %s", e.Status, e.Message)
}
