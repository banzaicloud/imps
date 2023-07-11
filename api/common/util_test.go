package common

import (
	"testing"

	"gotest.tools/assert"
)

func TestUtil_SecretNameFromURL(t *testing.T) {
	t.Parallel()
	type args struct {
		prefix string
		url    string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic functionality check",
			args: args{
				prefix: "test-prefix",
				url:    "testing.test",
			},
			want: "test-prefix-testing.test-pull-secret-4b3d6963",
		},
		{
			name: "invalid characters in URL",
			args: args{
				prefix: "test-prefix",
				url:    "+test'ing?.test!",
			},
			want: "test-prefix-test-ing-.test-pull-secret-fb0323c3",
		},
		{
			name: "URL is too long",
			args: args{
				prefix: "test-prefix",
				url:    "testing.loooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooonoooooooong.test",
			},
			want: "test-prefix-testing.loooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooonoooooooong.tes-pull-secret-240fd929",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			found := SecretNameFromURL(tt.args.prefix, tt.args.url)

			assert.Equal(t, tt.want, found)
		})
	}
}
