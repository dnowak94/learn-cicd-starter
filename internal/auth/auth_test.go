package auth

import (
	"errors"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input  http.Header
		errors error
		want   string
	}
	var authHeader string = "ApiKey aW4yOG1pbnV0ZXM6ZHVtbXk="
	splitAuthHeader := strings.Split(authHeader, " ")

	header1 := make(http.Header)
	header1.Add("Authorization", authHeader)
	header2 := make(http.Header)
	header2.Add("Authorization", splitAuthHeader[1])
	header3 := make(http.Header)

	tests := []test{
		{input: header1, errors: nil, want: splitAuthHeader[1]},
		{input: header2, errors: errors.New("malformed authorization header"), want: ""},
		{input: header3, errors: ErrNoAuthHeaderIncluded, want: ""},
	}

	for _, tc := range tests {

		got, err := GetAPIKey(tc.input)

		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
		if !reflect.DeepEqual(tc.errors, err) {
			t.Fatalf("expected: %v, got: %v", tc.errors, err)
		}
	}
}
