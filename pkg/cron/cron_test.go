package cron

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAsteriskCron(t *testing.T) {
	tests := []struct {
		Value                CronValue
		ExpectedParsedValues []int
	}{
		{
			Value: CronValue{
				Name:  "Test",
				Value: "*",
				Min:   10,
				Max:   15,
			},
			ExpectedParsedValues: []int{10, 11, 12, 13, 14, 15},
		},
	}
	for _, test := range tests {
		test.Value.parse()
		require.Equal(t, test.ExpectedParsedValues, test.Value.parsedValues)
	}
}
