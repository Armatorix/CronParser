package cron

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type Cron struct {
	Minute     CronValue
	Hour       CronValue
	DayOfMonth CronValue
	Month      CronValue
	DayOfWeek  CronValue
	Command    string
}

// For now return first occured error, can be extended to show all of them
func (c *Cron) parse() error {
	cronEntities := []*CronValue{
		&c.Minute,
		&c.Hour,
		&c.DayOfMonth,
		&c.Month,
		&c.DayOfWeek,
	}
	for _, cronEntity := range cronEntities {
		if err := cronEntity.parse(); err != nil {
			return errors.WithMessagef(err, "%s parsing failed", cronEntity.name)
		}
	}
	return nil
}

func (c Cron) String() string {
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%-14s %s\n", c.Minute, c.Hour, c.DayOfMonth, c.Month, c.DayOfWeek, "command", c.Command)
}

func NewFromOsArgs() (*Cron, error) {
	osArgs := os.Args
	if len(osArgs) != 2 {
		return nil, fmt.Errorf("incorrect number of command line arguments: %d", len(osArgs))
	}
	args := strings.Split(osArgs[1], " ")
	if len(args) != 6 {
		return nil, fmt.Errorf("incorrect number of cron arguments: %d", len(args))
	}

	c := &Cron{
		Minute: CronValue{
			name:  "minute",
			value: args[0],
			min:   0,
			max:   59,
		},
		Hour: CronValue{
			name:  "hour",
			value: args[1],
			min:   0,
			max:   23,
		},
		DayOfMonth: CronValue{
			name:  "day of month",
			value: args[2],
			min:   1,
			max:   31,
		},
		Month: CronValue{
			name:  "month",
			value: args[3],
			min:   1,
			max:   12,
		},
		DayOfWeek: CronValue{
			name:  "day of week",
			value: args[4],
			min:   0,
			max:   6,
		},
		Command: args[5],
	}
	err := c.parse()
	if err != nil {
		return nil, err
	}
	return c, nil
}
