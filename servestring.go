package http415

import (
	"net/http"
)

func ServeString(responseWriter http.ResponseWriter, value string, mediaTypes ...string) error {
	if nil == responseWriter {
		return ErrNilHTTPResponseWriter
	}

	{
		var accept string = acceptValue(mediaTypes...)
		if "" != accept {
			responseWriter.Header().Set("Accept", accept)
		}
	}

	responseWriter.WriteHeader(StatusCode)
	return nil
}
