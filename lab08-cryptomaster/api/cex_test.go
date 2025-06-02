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

func TestGetRate_InvalidCurrencyLength(t *testing.T) {
	_, err := api.GetRate("US") // only 2 letters

	if err == nil {
		t.Error("expected error for invalid currency length, got nil")
	}
}

func TestGetRate_UnsupportedCurrency(t *testing.T) {
	_, err := api.GetRate("XYZIOSD") // unlikely to be a valid currency

	if err == nil {
		t.Error("expected error for unsupported currency, got nil")
	}
}

func TestGetRate_ValidCurrency(t *testing.T) {
	rate, err := api.GetRate("BTC")

	if err != nil {
		t.Errorf("expected no error for valid currency, got %v", err)
	}

	if rate == nil || rate.Price == 0 {
		t.Errorf("expected valid rate data, got %+v", rate)
	}
}

func TestGetRate_TableDriven(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"empty string", "", true},
		{"too short", "EU", true},
		{"too long", "EURO", true},
		{"unsupported currency", "XYZBASDK", true},
		{"valid currency", "BTC", false}, // depends on live response
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rate, err := api.GetRate(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRate(%q) error = %v, wantErr = %v", tt.input, err, tt.wantErr)
			}

			if !tt.wantErr && (rate == nil || rate.Price == 0) {
				t.Errorf("expected valid rate for input %q, got %+v", tt.input, rate)
			}
		})
	}
}
