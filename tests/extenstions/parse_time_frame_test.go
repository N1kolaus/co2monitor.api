package extensions_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	ex "github.com/fminister/co2monitor.api/extensions"
)

func TestValidateTimeDuration(t *testing.T) {
	tests := []struct {
		input        string
		expectedDur  time.Duration
		expectErrMsg bool
	}{
		{"10m", 10 * time.Minute, false},
		{"5h", 5 * time.Hour, false},
		{"3d", 3 * 24 * time.Hour, false},
		{"abc", 6 * time.Hour, true},
		{"2w", 6 * time.Hour, true},
		{"2 h", 6 * time.Hour, true},
		{"20", 6 * time.Hour, true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			dur := ex.ValidateTimeDuration(test.input)

			assert.Equal(t, test.expectedDur, dur)
			assert.Equal(t, test.expectErrMsg, dur == 6*time.Hour)
		})
	}
}
