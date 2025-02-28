package auth

import (
	"net/http"
	"strings"
	"testing"
)

type authTest struct {
	headerKey   string
	headerValue string
	expected    string
	expectedErr string
}

var authTests = []authTest{
	{"", "", "", "no authorization header included"},
	{"Authorization", "", "", "no authorization header included"},
	{"Authorization", "Bearer 1234567890", "", "malformed authorization header"},
	{"Authorization", "ApiKey 1234567890", "1234567890", ""},
}

func TestGetAPIKey(t *testing.T) {

	var output string
	var err error
	header := make(http.Header)

	for _, test := range authTests {

		header.Set(test.headerKey, test.headerValue)
		output, err = GetAPIKey(header)
		if err != nil {
			if !strings.Contains(err.Error(), test.expectedErr) {
				t.Errorf("Error was not what we expected. test.expectedErr = %s, err.Error() = %s", test.expectedErr, err.Error())
				return
			}
		}
		if output != test.expected {
			t.Errorf("Output was not what we expected. test.expected = %s, output = %s", test.expected, output)
			return
		}
	}
}
