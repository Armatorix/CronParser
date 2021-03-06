package cron

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewFromOsArgsCornerCases(t *testing.T) {
	tests := []struct {
		name         string
		osArgs       []string
		errSubstring string
	}{
		{
			name:         "no arguments",
			errSubstring: "incorrect number of command line arguments",
			osArgs:       []string{"cmdName"},
		},
		{
			name:         "incorrect amount of values in cron argument",
			errSubstring: "incorrect number of cron argument",
			osArgs:       []string{"cmdName", "1 2 3 4"},
		},
		{
			name:         "bad arguments",
			errSubstring: "day of week parsing failed",
			osArgs:       []string{"cmdName", "1 2 3 4 test cmd"},
		},
	}

	for _, test := range tests {
		os.Args = test.osArgs
		c, err := NewFromOsArgs()
		require.Error(t, err)
		require.Nil(t, c)
		require.Contains(t, err.Error(), test.errSubstring, test)
	}
}
func TestNewFromOsArgs(t *testing.T) {
	tests := []struct {
		name     string
		osArgs   []string
		expected string
	}{
		{
			name:   "example1",
			osArgs: []string{"cmd", "1 2 3 4 5 /usr/bin/yes"},
			expected: `minute         1
hour           2
day of month   3
month          4
day of week    5
command        /usr/bin/yes
`,
		},
		{
			name:   "example2",
			osArgs: []string{"cmd", "*/15 0 1,15 * 1-5 /usr/bin/find"},
			expected: `minute         0 15 30 45
hour           0
day of month   1 15
month          1 2 3 4 5 6 7 8 9 10 11 12
day of week    1 2 3 4 5
command        /usr/bin/find
`,
		},
		{
			name:   "all cmd",
			osArgs: []string{"cmd", "* * * * * cmd"},
			expected: `minute         0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59
hour           0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23
day of month   1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31
month          1 2 3 4 5 6 7 8 9 10 11 12
day of week    0 1 2 3 4 5 6
command        cmd
`,
		},
	}

	for _, test := range tests {
		os.Args = test.osArgs
		c, err := NewFromOsArgs()
		require.NoError(t, err)
		require.NotNil(t, c)
		require.Equal(t, test.expected, c.String())
	}
}
