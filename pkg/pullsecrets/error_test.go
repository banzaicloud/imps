package pullsecrets

import (
	"testing"

	"emperror.dev/errors"
	"gotest.tools/assert"
)

func TestError_NewErrorsPerSecret(t *testing.T) {

	tests := []struct {
		name string
		want ErrorsPerSecret
	}{
		{
			name: "basic functionality test",
			want: ErrorsPerSecret{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := NewErrorsPerSecret()

			assert.DeepEqual(t, tt.want, found)
		})
	}
}

func TestError_AddSecret(t *testing.T) {
	type args struct {
		name string
	}

	tests := []struct {
		name            string
		args            args
		errorsPerSecret ErrorsPerSecret
	}{
		{
			name: "basic functionality test",
			args: args{
				name: "testSecret",
			},
			errorsPerSecret: ErrorsPerSecret{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.errorsPerSecret.AddSecret(tt.args.name)

			assert.Assert(t, tt.errorsPerSecret["testSecret"] == nil)
		})
	}
}

func TestError_SetSecretError(t *testing.T) {
	type args struct {
		name string
		err  error
	}

	tests := []struct {
		name                 string
		args                 args
		errorsPerSecret      ErrorsPerSecret
		expectedErrorMessage string
	}{
		{
			name: "empty error message",
			args: args{
				name: "testSecret",
				err:  errors.New(""),
			},
			errorsPerSecret:      ErrorsPerSecret{},
			expectedErrorMessage: "",
		},
		{
			name: "non-empty error message",
			args: args{
				name: "testSecret",
				err:  errors.New("testError"),
			},
			errorsPerSecret:      ErrorsPerSecret{},
			expectedErrorMessage: "testError",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.errorsPerSecret.SetSecretError(tt.args.name, tt.args.err)

			assert.Assert(t, tt.errorsPerSecret["testSecret"].Error() == tt.expectedErrorMessage)
		})
	}
}

func TestError_AsStatus(t *testing.T) {

	testErrorPerSecret := ErrorsPerSecret{}
	testErrorPerSecret.SetSecretError("testSecret", errors.New("testError"))
	testErrorPerSecret.AddSecret("testSecret2")

	tests := []struct {
		name            string
		errorsPerSecret ErrorsPerSecret
		wanted          map[string]string
	}{
		{
			name:            "basic functionality test",
			errorsPerSecret: testErrorPerSecret,
			wanted: map[string]string{
				"testSecret":  "testError",
				"testSecret2": "Ok",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := tt.errorsPerSecret.AsStatus()

			assert.DeepEqual(t, tt.wanted, found)
		})
	}
}

func TestError_FailedSecrets(t *testing.T) {

	testErrorPerSecret := ErrorsPerSecret{}
	testErrorPerSecret.SetSecretError("testSecret", errors.New("testError"))
	testErrorPerSecret.AddSecret("testSecret2")

	tests := []struct {
		name            string
		errorsPerSecret ErrorsPerSecret
		wanted          []string
	}{
		{
			name:            "basic functionality test",
			errorsPerSecret: testErrorPerSecret,
			wanted:          []string{"testSecret"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := tt.errorsPerSecret.FailedSecrets()

			assert.DeepEqual(t, tt.wanted, found)
		})
	}
}

func TestError_AsError(t *testing.T) {

	testErrorPerSecret := ErrorsPerSecret{}
	testErrorPerSecret.SetSecretError("testSecret", errors.New("testError"))
	testErrorPerSecret.AddSecret("testSecret2")

	tests := []struct {
		name            string
		errorsPerSecret ErrorsPerSecret
		wanted          error
	}{
		{
			name:            "no invalid secrets",
			errorsPerSecret: ErrorsPerSecret{},
			wanted:          nil,
		},
		{
			name:            "one invalid secret",
			errorsPerSecret: testErrorPerSecret,
			wanted:          errors.NewWithDetails("some source secrets failed to render", "failed_secrets", "testSecret"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := tt.errorsPerSecret.AsError()

			if len(tt.errorsPerSecret.FailedSecrets()) > 0 {
				assert.Equal(t, tt.wanted.Error(), found.Error())
			} else {
				assert.Equal(t, tt.wanted, found)
			}
		})
	}
}
