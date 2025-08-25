package auth

import (
	"net/http"
	"strings"
	"testing"
)

func assertLen(t *testing.T, val string, length int) {
	if len(val) != length {
		t.Errorf("len %s != %d", val, length)
	}
}

func assertEqual(t *testing.T, val1 string, val2 string) {
	if strings.Compare(val1, val2) != 0 {
		t.Errorf("%s doesnt equal to %s", val1, val2)
	}
}

func assertError(t *testing.T, val error) {
	if val == nil {
		t.Fatalf("Didn't get error")
	}
}

func requireNoError(t *testing.T, val error) {
	if val != nil {
		t.Fatalf("Error doesn't equal nil: %s", val.Error())
	}
}

func TestGetApiKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("AuthorizationZSJDIUSAD", "sdjoasd")
	str, err := GetAPIKey(headers)
	assertError(t, err)
	assertLen(t, str, 0)

	headers = http.Header{}
	headers.Set("Authorization", "apikey YOLO")
	str, err = GetAPIKey(headers)
	assertError(t, err)
	assertLen(t, str, 0)

	headers = http.Header{}
	headers.Set("Authorization", "ApiKey")
	str, err = GetAPIKey(headers)
	assertError(t, err)
	assertLen(t, str, 0)

	headers = http.Header{}
	headers.Set("Authorization", "ApiKey Hello World!")
	str, err = GetAPIKey(headers)
	assertError(t, err)
	assertLen(t, str, 0)
	//requireNoError(t, err)
	//assertEqual(t, str, "Hello")
}
