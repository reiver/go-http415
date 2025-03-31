package http415

import (
	"net/http"
)

func Serve(responseWriter http.ResponseWriter, mediaTypes ...string) error {
	return ServeString(responseWriter, DefaultStatusText, mediaTypes...)
}
