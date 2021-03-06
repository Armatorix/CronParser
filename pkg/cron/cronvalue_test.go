package cron

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAsteriskCron(t *testing.T) {
	tests := []struct {
		value                CronValue
		expectedParsedValues []int64
	}{
		{
			value: CronValue{
				name:  "Test dummy asterisk",
				value: "*",
				min:   10,
				max:   15,
			},
			expectedParsedValues: []int64{10, 11, 12, 13, 14, 15},
		},
		{
			value: CronValue{
				name:  "Test week days asterisk",
				value: "*",
				min:   0,
				max:   6,
			},
			expectedParsedValues: []int64{0, 1, 2, 3, 4, 5, 6},
		},
		{
			value: CronValue{
				name:  "Test mixed asterisk",
				value: "1,*,5,2-4",
				min:   0,
				max:   6,
			},
			expectedParsedValues: []int64{0, 1, 2, 3, 4, 5, 6},
		},
	}
	for _, test := range tests {
		require.NoError(t, test.value.parse())
		require.Equal(t, test.expectedParsedValues, test.value.parsedValues, test)
	}
}

func TestStepCron(t *testing.T) {
	tests := []struct {
		value                CronValue
		expectedParsedValues []int64
	}{
		{
			value: CronValue{
				name:  "Test asterisk with step",
				value: "*/2",
				min:   0,
				max:   7,
			},
			expectedParsedValues: []int64{0, 2, 4, 6},
		},
		{
			value: CronValue{
				name:  "Test asterisk with step",
				value: "*/3",
				min:   0,
				max:   7,
			},
			expectedParsedValues: []int64{0, 3, 6},
		},
		{
			value: CronValue{
				name:  "Test asterisk with step",
				value: "*/6",
				min:   0,
				max:   7,
			},
			expectedParsedValues: []int64{0, 6},
		},
	}
	for _, test := range tests {
		require.NoError(t, test.value.parse())
		require.Equal(t, test.expectedParsedValues, test.value.parsedValues, test)
	}
}

func TestRangeCron(t *testing.T) {
	tests := []struct {
		value                CronValue
		expectedParsedValues []int64
	}{
		{
			value: CronValue{
				name:  "Test month day range",
				value: "10-15",
				min:   1,
				max:   31,
			},
			expectedParsedValues: []int64{10, 11, 12, 13, 14, 15},
		},
		{
			value: CronValue{
				name:  "Test whole week of week days",
				value: "0-6",
				min:   0,
				max:   6,
			},
			expectedParsedValues: []int64{0, 1, 2, 3, 4, 5, 6},
		},
		{
			value: CronValue{
				name:  "Test ranges next to each other",
				value: "3-4,5-6,7-8",
				min:   0,
				max:   30,
			},
			expectedParsedValues: []int64{3, 4, 5, 6, 7, 8},
		},
		{
			value: CronValue{
				name:  "Test ranges with similar part",
				value: "3-6,5-7",
				min:   0,
				max:   30,
			},
			expectedParsedValues: []int64{3, 4, 5, 6, 7},
		},
	}
	for _, test := range tests {
		require.NoError(t, test.value.parse())
		require.Equal(t, test.expectedParsedValues, test.value.parsedValues, test)
	}
}

func TestSingleValueCron(t *testing.T) {
	tests := []struct {
		value                CronValue
		expectedParsedValues []int64
	}{
		{
			value: CronValue{
				name:  "Test single day from month days",
				value: "5",
				min:   1,
				max:   31,
			},
			expectedParsedValues: []int64{5},
		},
		{
			value: CronValue{
				name:  "Test multiple days from month days",
				value: "5,6,7,8,5,2",
				min:   1,
				max:   31,
			},
			expectedParsedValues: []int64{2, 5, 6, 7, 8},
		},
	}
	for _, test := range tests {
		require.NoError(t, test.value.parse())
		require.Equal(t, test.expectedParsedValues, test.value.parsedValues, test)
	}
}

func TestMixedCronTypes(t *testing.T) {
	tests := []struct {
		value                CronValue
		expectedParsedValues []int64
	}{
		{
			value: CronValue{
				name:  "Test single value with range excludive",
				value: "5,6-8",
				min:   1,
				max:   31,
			},
			expectedParsedValues: []int64{5, 6, 7, 8},
		},
		{
			value: CronValue{
				name:  "Test single value with range includive",
				value: "8,6-10",
				min:   1,
				max:   31,
			},
			expectedParsedValues: []int64{6, 7, 8, 9, 10},
		},
		{
			value: CronValue{
				name:  "Test single value, range and asterisk",
				value: "2,4-5,*",
				min:   0,
				max:   6,
			},
			expectedParsedValues: []int64{0, 1, 2, 3, 4, 5, 6},
		},
	}
	for _, test := range tests {
		require.NoError(t, test.value.parse())
		require.Equal(t, test.expectedParsedValues, test.value.parsedValues, test)
	}
}
