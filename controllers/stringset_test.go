package controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringSet_Has(t *testing.T) {
	type args struct {
		needle string
	}

	haystack := StringSet{"testing1", "testing2", "testing3", "testing4", "testing5", "testing6", "testing7", "testing8"}

	tests := []struct {
		name string
		set  StringSet
		args args
		want bool
	}{
		{
			name: "positive",
			set:  haystack,
			args: args{
				needle: "testing3",
			},
			want: true,
		},
		{
			name: "negative",
			set:  haystack,
			args: args{
				needle: "testing9",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := tt.set.Has(tt.args.needle)

			assert.Equal(t, tt.want, found)
		})
	}
}
