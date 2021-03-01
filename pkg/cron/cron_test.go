package cron

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAsteriskCron(t *testing.T) {
	tests := []struct {
		Value                CronValue
		ExpectedParsedValues []int64
	}{
		{
			Value: CronValue{
				Name:  "Test dummy asterisk",
				Value: "*",
				Min:   10,
				Max:   15,
			},
			ExpectedParsedValues: []int64{10, 11, 12, 13, 14, 15},
		},
		{
			Value: CronValue{
				Name:  "Test week days asterisk",
				Value: "*",
				Min:   0,
				Max:   6,
			},
			ExpectedParsedValues: []int64{0, 1, 2, 3, 4, 5, 6},
		},
		{
			Value: CronValue{
				Name:  "Test mixed asterisk",
				Value: "21,*,37,12-15",
				Min:   0,
				Max:   6,
			},
			ExpectedParsedValues: []int64{0, 1, 2, 3, 4, 5, 6},
		},
	}
	for _, test := range tests {
		test.Value.parse()
		require.Equal(t, test.ExpectedParsedValues, test.Value.parsedValues, test)
	}
}

func TestStepCron(t *testing.T) {
	tests := []struct {
		Value                CronValue
		ExpectedParsedValues []int64
	}{
		{
			Value: CronValue{
				Name:  "Test asterisk with step",
				Value: "*/2",
				Min:   0,
				Max:   7,
			},
			ExpectedParsedValues: []int64{0, 2, 4, 6},
		},
		{
			Value: CronValue{
				Name:  "Test asterisk with step",
				Value: "*/3",
				Min:   0,
				Max:   7,
			},
			ExpectedParsedValues: []int64{0, 3, 6},
		},
		{
			Value: CronValue{
				Name:  "Test asterisk with step",
				Value: "*/6",
				Min:   0,
				Max:   7,
			},
			ExpectedParsedValues: []int64{0, 6},
		},
	}
	for _, test := range tests {
		test.Value.parse()
		require.Equal(t, test.ExpectedParsedValues, test.Value.parsedValues, test)
	}
}

func TestRangeCron(t *testing.T) {
	tests := []struct {
		Value                CronValue
		ExpectedParsedValues []int64
	}{
		{
			Value: CronValue{
				Name:  "Test month day range",
				Value: "10-15",
				Min:   1,
				Max:   31,
			},
			ExpectedParsedValues: []int64{10, 11, 12, 13, 14, 15},
		},
		{
			Value: CronValue{
				Name:  "Test whole week of week days",
				Value: "0-6",
				Min:   0,
				Max:   6,
			},
			ExpectedParsedValues: []int64{0, 1, 2, 3, 4, 5, 6},
		},
		{
			Value: CronValue{
				Name:  "Test ranges next to each other",
				Value: "3-4,5-6,7-8",
				Min:   0,
				Max:   30,
			},
			ExpectedParsedValues: []int64{3, 4, 5, 6, 7, 8},
		},
		{
			Value: CronValue{
				Name:  "Test ranges with similar part",
				Value: "3-6,5-7",
				Min:   0,
				Max:   30,
			},
			ExpectedParsedValues: []int64{3, 4, 5, 6, 7},
		},
	}
	for _, test := range tests {
		test.Value.parse()
		require.Equal(t, test.ExpectedParsedValues, test.Value.parsedValues, test)
	}
}

func TestSingleValueCron(t *testing.T) {
	tests := []struct {
		Value                CronValue
		ExpectedParsedValues []int64
	}{
		{
			Value: CronValue{
				Name:  "Test single day from month days",
				Value: "5",
				Min:   1,
				Max:   31,
			},
			ExpectedParsedValues: []int64{5},
		},
		{
			Value: CronValue{
				Name:  "Test multiple days from month days",
				Value: "5,6,7,8,5,2",
				Min:   1,
				Max:   31,
			},
			ExpectedParsedValues: []int64{2, 5, 6, 7, 8},
		},
	}
	for _, test := range tests {
		test.Value.parse()
		require.Equal(t, test.ExpectedParsedValues, test.Value.parsedValues, test)
	}
}

func TestMixedCronTypes(t *testing.T) {
	tests := []struct {
		Value                CronValue
		ExpectedParsedValues []int64
	}{
		{
			Value: CronValue{
				Name:  "Test single value with range excludive",
				Value: "5,6-8",
				Min:   1,
				Max:   31,
			},
			ExpectedParsedValues: []int64{5, 6, 7, 8},
		},
		{
			Value: CronValue{
				Name:  "Test single value with range includive",
				Value: "8,6-10",
				Min:   1,
				Max:   31,
			},
			ExpectedParsedValues: []int64{6, 7, 8, 9, 10},
		},
		{
			Value: CronValue{
				Name:  "Test single value, range and asterisk",
				Value: "2,4-5,*",
				Min:   0,
				Max:   6,
			},
			ExpectedParsedValues: []int64{0, 1, 2, 3, 4, 5},
		},
	}
	for _, test := range tests {
		test.Value.parse()
		require.Equal(t, test.ExpectedParsedValues, test.Value.parsedValues, test)
	}
}
