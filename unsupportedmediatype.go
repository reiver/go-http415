package http415

import (
	"net/http"
)

func UnsupportedMediaType(responseWriter http.ResponseWriter, request *http.Request) {
	Serve(responseWriter)
}
