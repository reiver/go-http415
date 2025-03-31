package http415

import (
	"fmt"
)

func acceptValue(mediaTypes ...string) string {

	var value string

	for _, mediaType := range mediaTypes {
		if "" == mediaType {
			continue
		}

		switch value {
		case "":
			value = mediaType
		default:
			value = fmt.Sprintf("%s, %s", value, mediaType)
		}
	}

	return value
}
