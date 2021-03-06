package cron

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAsteriskCron(t *testing.T) {
	tests := []struct {
		name           string
		cronValue      CronValue
		expectedValues []int64
		expectedString string
	}{
		{
			name: "test dummy asterisk",
			cronValue: CronValue{
				name:  "test",
				value: "*",
				min:   10,
				max:   15,
			},
			expectedValues: []int64{10, 11, 12, 13, 14, 15},
			expectedString: "test           10 11 12 13 14 15",
		},
		{
			name: "test week days asterisk",
			cronValue: CronValue{
				name:  "test",
				value: "*",
				min:   0,
				max:   6,
			},
			expectedValues: []int64{0, 1, 2, 3, 4, 5, 6},
			expectedString: "test           0 1 2 3 4 5 6",
		},
		{
			name: "test mixed asterisk",
			cronValue: CronValue{
				name:  "test",
				value: "1,*,5,2-4",
				min:   0,
				max:   6,
			},
			expectedValues: []int64{0, 1, 2, 3, 4, 5, 6},
			expectedString: "test           0 1 2 3 4 5 6",
		},
	}
	for _, test := range tests {
		require.NoError(t, test.cronValue.parse())
		require.Equal(t, test.expectedValues, test.cronValue.parsedValues, test)
		require.Equal(t, test.expectedString, test.cronValue.String())
	}
}

func TestStepCron(t *testing.T) {
	tests := []struct {
		name           string
		cronValue      CronValue
		expectedValues []int64
		expectedString string
	}{
		{
			name: "test asterisk with step",
			cronValue: CronValue{
				name:  "test",
				value: "*/2",
				min:   0,
				max:   7,
			},
			expectedValues: []int64{0, 2, 4, 6},
			expectedString: "test           0 2 4 6",
		},
		{
			name: "test asterisk with step",
			cronValue: CronValue{
				name:  "test",
				value: "*/3",
				min:   0,
				max:   7,
			},
			expectedValues: []int64{0, 3, 6},
			expectedString: "test           0 3 6",
		},
		{
			name: "test asterisk with step",
			cronValue: CronValue{
				name:  "test",
				value: "*/6",
				min:   0,
				max:   7,
			},
			expectedValues: []int64{0, 6},
			expectedString: "test           0 6",
		},
	}
	for _, test := range tests {
		require.NoError(t, test.cronValue.parse())
		require.Equal(t, test.expectedValues, test.cronValue.parsedValues, test)
		require.Equal(t, test.expectedString, test.cronValue.String())
	}
}

func TestRangeCron(t *testing.T) {
	tests := []struct {
		name           string
		cronValue      CronValue
		expectedValues []int64
		expectedString string
	}{
		{
			name: "test month day range",
			cronValue: CronValue{
				name:  "test",
				value: "10-15",
				min:   1,
				max:   31,
			},
			expectedValues: []int64{10, 11, 12, 13, 14, 15},
			expectedString: "test           10 11 12 13 14 15",
		},
		{
			name: "test whole week of week days",
			cronValue: CronValue{
				name:  "test",
				value: "0-6",
				min:   0,
				max:   6,
			},
			expectedValues: []int64{0, 1, 2, 3, 4, 5, 6},
			expectedString: "test           0 1 2 3 4 5 6",
		},
		{
			name: "test ranges next to each other",
			cronValue: CronValue{
				name:  "test",
				value: "3-4,5-6,7-8",
				min:   0,
				max:   30,
			},
			expectedValues: []int64{3, 4, 5, 6, 7, 8},
			expectedString: "test           3 4 5 6 7 8",
		},
		{
			name: "test ranges with similar part",
			cronValue: CronValue{
				name:  "test",
				value: "3-6,5-7",
				min:   0,
				max:   30,
			},
			expectedValues: []int64{3, 4, 5, 6, 7},
			expectedString: "test           3 4 5 6 7",
		},
	}
	for _, test := range tests {
		require.NoError(t, test.cronValue.parse())
		require.Equal(t, test.expectedValues, test.cronValue.parsedValues, test)
		require.Equal(t, test.expectedString, test.cronValue.String())
	}
}

func TestSingleValueCron(t *testing.T) {
	tests := []struct {
		name           string
		cronValue      CronValue
		expectedValues []int64
		expectedString string
	}{
		{
			name: "test single day from month days",
			cronValue: CronValue{
				name:  "test",
				value: "5",
				min:   1,
				max:   31,
			},
			expectedValues: []int64{5},
			expectedString: "test           5",
		},
		{
			name: "test multiple days from month days",
			cronValue: CronValue{
				name:  "test",
				value: "5,6,7,8,5,2",
				min:   1,
				max:   31,
			},
			expectedValues: []int64{2, 5, 6, 7, 8},
			expectedString: "test           2 5 6 7 8",
		},
	}
	for _, test := range tests {
		require.NoError(t, test.cronValue.parse())
		require.Equal(t, test.expectedValues, test.cronValue.parsedValues, test)
		require.Equal(t, test.expectedString, test.cronValue.String())
	}
}

func TestMixedCronTypes(t *testing.T) {
	tests := []struct {
		name           string
		cronValue      CronValue
		expectedValues []int64
		expectedString string
	}{
		{
			name: "test single value with range excludive",
			cronValue: CronValue{
				name:  "test",
				value: "5,6-8",
				min:   1,
				max:   31,
			},
			expectedValues: []int64{5, 6, 7, 8},
			expectedString: "test           5 6 7 8",
		},
		{
			name: "test single value with range includive",
			cronValue: CronValue{
				name:  "test",
				value: "8,6-10",
				min:   1,
				max:   31,
			},
			expectedValues: []int64{6, 7, 8, 9, 10},
			expectedString: "test           6 7 8 9 10",
		},
		{
			name: "test single value, range and asterisk",
			cronValue: CronValue{
				name:  "test",
				value: "2,4-5,*",
				min:   0,
				max:   6,
			},
			expectedValues: []int64{0, 1, 2, 3, 4, 5, 6},
			expectedString: "test           0 1 2 3 4 5 6",
		},
	}
	for _, test := range tests {
		require.NoError(t, test.cronValue.parse())
		require.Equal(t, test.expectedValues, test.cronValue.parsedValues, test)
		require.Equal(t, test.expectedString, test.cronValue.String())
	}
}
