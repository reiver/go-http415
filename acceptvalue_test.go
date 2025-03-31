package http415

import (
	"testing"
)

func TestAcceptValue(t *testing.T) {

	tests := []struct{
		MediaTypes []string
		Expected string
	}{
		{

		},



		{
			MediaTypes: []string{"application/activity+json"},
			Expected:            `application/activity+json`,
		},
		{
			MediaTypes: []string{"application/activity+json", `application/ld+json; profile="https://www.w3.org/ns/activitystreams"`},
			Expected:            `application/activity+json, application/ld+json; profile="https://www.w3.org/ns/activitystreams"`,
		},
	}

	for testNumber, test := range tests {

		actual := acceptValue(test.MediaTypes...)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual HTTP response Accept value is not what was expected.", testNumber)
			t.Logf("EXPECTED:...\n%s", expected)
			t.Logf("ACTUAL:...\n%s", actual)
			t.Logf("MEDIA-TYPES: %#v", test.MediaTypes)
			continue
		}
	}
}
