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

func TestApplyNumber(t *testing.T) {
	min, max := int64(10), int64(20)
	tests := []struct {
		name    string
		numbers []int64
		err     error
	}{
		{
			name:    "min edge",
			numbers: []int64{min},
			err:     nil,
		},
		{
			name:    "max edge",
			numbers: []int64{max},
			err:     nil,
		},
		{
			name:    "values from range",
			numbers: []int64{10, 20},
			err:     nil,
		},
		{
			name:    "values out of range",
			numbers: []int64{-10, 0, 9, 21, 37},
			err:     errOutOfBound,
		},
	}

	for _, test := range tests {
		ex, err := New(min, max)
		require.NoError(t, err)
		if test.err == nil {
			for i, v := range test.numbers {
				err = ex.ApplyNumber(v)
				require.NoError(t, err)
				require.Equal(t, test.numbers[:i+1], ex.ToInt64Slice())
			}
		} else {
			for _, v := range test.numbers {
				err = ex.ApplyNumber(v)
				require.ErrorIs(t, err, test.err)
			}
		}
	}
}

func TestApplySlice(t *testing.T) {
	min, max := int64(10), int64(20)
	tests := []struct {
		name    string
		numbers []int64
		err     error
	}{
		{
			name:    "values from range",
			numbers: []int64{10, 11, 12, 15, 17, 20},
			err:     nil,
		},
		{
			name:    "all values",
			numbers: []int64{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			err:     nil,
		},
		{
			name:    "values out of range",
			numbers: []int64{-10, 0, 9, 21, 37},
			err:     errOutOfBound,
		},
		{
			name:    "mixed from and out of range",
			numbers: []int64{0, 10, 20, 30},
			err:     errOutOfBound,
		},
	}

	for _, test := range tests {
		ex, err := New(min, max)
		require.NoError(t, err)
		err = ex.ApplySlice(test.numbers)
		if test.err == nil {
			require.NoError(t, err)
			require.Equal(t, test.numbers, ex.ToInt64Slice())
		} else {
			require.ErrorIs(t, err, test.err)
		}
	}
}
