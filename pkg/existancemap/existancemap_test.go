package existancemap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		min      int64
		max      int64
		err      error
		sliceLen int
	}{
		{
			name:     "one value 0",
			min:      0,
			max:      0,
			err:      nil,
			sliceLen: 1,
		},
		{
			name:     "from 0 to 10",
			min:      0,
			max:      10,
			err:      nil,
			sliceLen: 11,
		},
		{
			name:     "from 10 to 0",
			min:      10,
			max:      0,
			err:      errMinGTMax,
			sliceLen: 0,
		},
		{
			name:     "from -10 to 10",
			min:      -10,
			max:      10,
			err:      nil,
			sliceLen: 21,
		},
	}

	for _, test := range tests {
		ex, err := New(test.min, test.max)
		require.ErrorIs(t, test.err, err)
		if test.err == nil {
			require.Equal(t, test.min, ex.min)
			require.Equal(t, test.max, ex.max)
			require.Equal(t, test.sliceLen, len(ex.existance))
			require.Equal(t, test.min, ex.min)
		}
	}
}
