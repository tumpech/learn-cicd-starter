package auth

import (
	"errors"
	"net/http"
	"testing"
)

type authTest struct {
	header   string
	expected string
	err      error
}

var authTests = []authTest{
	{"", "", ErrNoAuthHeaderIncluded},
	{"Bearer 1234567890", "", errors.New("malformed authorization header")},
	{"ApiKey 1234567890", "1234567890", nil},
}

func TestGetAPIKey(t *testing.T) {

	var output string
	var err error
	header := make(http.Header)

	for _, test := range authTests {

		header.Set("Authorization", test.header)
		output, err = GetAPIKey(header)
		//if !errors.Is(err, test.err) {
		//	t.Errorf("Error was not what we expected. test.err = %s, err = %s", test.err, err)
		//}
		if output != test.expected {
			t.Errorf("Output was not what we expected. test.expected = %s, output = %s", test.err, err)
		}
	}
}
