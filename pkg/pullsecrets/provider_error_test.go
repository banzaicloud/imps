package pullsecrets

import (
	"testing"

	"emperror.dev/errors"
	"gotest.tools/assert"
)

func TestErroredCredentialProvider_NewErroredCredentialProvider(t *testing.T) {
	type args struct {
		err error
	}

	tests := []struct {
		name   string
		args   args
		wanted ErroredCredentialProvider
	}{
		{
			name: "empty error",
			args: args{
				err: nil,
			},
			wanted: ErroredCredentialProvider{
				Error: nil,
			},
		},
		{
			name: "non-empty error",
			args: args{
				err: errors.New("testError"),
			},
			wanted: ErroredCredentialProvider{
				Error: errors.New("testError"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := NewErroredCredentialProvider(tt.args.err)

			if tt.wanted.Error != nil {
				assert.Equal(t, tt.wanted.Error.Error(), found.Error.Error())
			} else {
				assert.DeepEqual(t, tt.wanted, found)
			}
		})
	}
}
