package v1alpha1

import (
	"testing"

	"gotest.tools/assert"
)

func TestLabelSet_DeepCopy(t *testing.T) {
	t.Parallel()
	testlabelSet := map[string]string{
		"testLabel": "true",
	}

	tests := []struct {
		name     string
		labelSet LabelSet
	}{
		{
			name:     "deep copy test",
			labelSet: testlabelSet,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			labelSetCopy := tt.labelSet.DeepCopy()

			assert.DeepEqual(t, tt.labelSet, labelSetCopy)
			assert.Assert(t, &tt.labelSet != &labelSetCopy)
		})
	}
}
