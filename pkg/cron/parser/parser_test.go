package parser

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseRange(t *testing.T) {
	tests := []struct {
		rangeStr string
		min      int64
		max      int64
		err      error
	}{
		{
			rangeStr: "",
			min:      0,
			max:      0,
			err:      errWrongFormat,
		},
		{
			rangeStr: "-",
			min:      0,
			max:      0,
			err:      strconv.ErrSyntax,
		},
		{
			rangeStr: "-0",
			min:      0,
			max:      0,
			err:      strconv.ErrSyntax,
		},
		{
			rangeStr: "0-",
			min:      0,
			max:      0,
			err:      strconv.ErrSyntax,
		},
		{
			rangeStr: "1-10",
			min:      1,
			max:      10,
			err:      nil,
		},
		{
			rangeStr: "10-1",
			min:      10,
			max:      1,
			err:      errMinGTMax,
		},
	}
	for _, test := range tests {
		min, max, err := ParseRange(test.rangeStr)
		require.Equal(t, test.min, min)
		require.Equal(t, test.max, max)
		require.ErrorIs(t, err, test.err, test)
	}
}

func TestParseStep(t *testing.T) {
	tests := []struct {
		stepStr  string
		min      int64
		max      int64
		err      error
		expected []int64
	}{
		{
			stepStr:  "*/",
			min:      0,
			max:      5,
			err:      errWrongFormat,
			expected: nil,
		},
		{
			stepStr:  "/",
			min:      0,
			max:      5,
			err:      errWrongFormat,
			expected: nil,
		},
		{
			stepStr:  "*//",
			min:      0,
			max:      5,
			err:      strconv.ErrSyntax,
			expected: nil,
		},
		{
			stepStr:  "*/6",
			min:      0,
			max:      5,
			err:      errStepTooBig,
			expected: nil,
		},
		{
			stepStr:  "*/5",
			min:      0,
			max:      5,
			err:      nil,
			expected: []int64{0, 5},
		},
		{
			stepStr:  "*/2",
			min:      0,
			max:      5,
			err:      nil,
			expected: []int64{0, 2, 4},
		},
	}
	for _, test := range tests {
		vals, err := ParseStep(test.stepStr, test.min, test.max)
		require.ErrorIs(t, err, test.err, test)
		require.Equal(t, test.expected, vals)
	}
}
