package log

import (
	"testing"

	"emperror.dev/errors"
	"gotest.tools/assert"
)

func TestConfig_Validate(t *testing.T) {

	tests := []struct {
		name        string
		config      Config
		want        Config
		expectedErr error
	}{
		{
			name:   "format string is empty",
			config: Config{},
			want: Config{
				Format: "logfmt",
			},
			expectedErr: nil,
		},
		{
			name: "format string is good",
			config: Config{
				Format: "json",
			},
			want: Config{
				Format: "json",
			},
			expectedErr: nil,
		},
		{
			name: "format string is wrong",
			config: Config{
				Format: "wrongFormat",
			},
			want: Config{
				Format: "wrongFormat",
			},
			expectedErr: errors.New("invalid log format: wrongFormat"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, err := tt.config.Validate()

			assert.Equal(t, tt.want, found)
			if tt.expectedErr != nil {
				assert.Error(t, err, tt.expectedErr.Error())
			}
		})
	}
}
