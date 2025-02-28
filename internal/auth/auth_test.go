package auth

import (
	"net/http"
	"strings"
	"testing"
)

type authTest struct {
	header      string
	expected    string
	expectedErr string
}

var authTests = []authTest{
	{"", "", ErrNoAuthHeaderIncluded.Error()},
	{"Bearer 1234567890", "", "malformed authorization header"},
	{"ApiKey 1234567890", "1234567890", ""},
}

func TestGetAPIKey(t *testing.T) {

	var output string
	var err error
	header := make(http.Header)

	for _, test := range authTests {

		header.Set("Authorization", test.header)
		output, err = GetAPIKey(header)
		if err != nil {
			if !strings.Contains(err.Error(), test.expectedErr) {
				t.Errorf("Error was not what we expected. test.err = %s, err = %s", test.expectedErr, err.Error())
				return
			}
		}
		if output != test.expected {
			t.Errorf("Output was not what we expected. test.expected = %s, output = %s", test.expected, output)
			return
		}
	}
}
