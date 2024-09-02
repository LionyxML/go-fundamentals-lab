package api_test

import (
	"cryptomaster/api"
	"testing"
)

func TestApiCall(t *testing.T) {
	_, err := api.GetRate("")

	if err == nil {
		t.Error("Error was not found")
	}
}
