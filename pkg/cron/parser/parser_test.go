package parser

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseRange(t *testing.T) {
	tests := []struct {
		strRange string
		min      int64
		max      int64
		err      error
	}{
		{
			strRange: "",
			min:      0,
			max:      0,
			err:      errWrongFormat,
		},
		{
			strRange: "-",
			min:      0,
			max:      0,
			err:      strconv.ErrSyntax,
		},
		{
			strRange: "-0",
			min:      0,
			max:      0,
			err:      strconv.ErrSyntax,
		},
		{
			strRange: "0-",
			min:      0,
			max:      0,
			err:      strconv.ErrSyntax,
		},
		{
			strRange: "1-10",
			min:      1,
			max:      10,
			err:      nil,
		},
		{
			strRange: "10-1",
			min:      10,
			max:      1,
			err:      errMinGTMax,
		},
	}
	for _, test := range tests {
		min, max, err := ParseRange(test.strRange)
		require.Equal(t, test.min, min)
		require.Equal(t, test.max, max)
		require.ErrorIs(t, err, test.err, test)
	}
}
