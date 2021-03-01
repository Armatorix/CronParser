package cron

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// CronValue provides minimal entity from cron object
// it handles minute, hour, day, month day of month
type CronValue struct {
	Name  string
	Value string
	Min   int64
	Max   int64

	parsedValues []int64
}

func (c CronValue) String() string {
	values := fmt.Sprint(c.parsedValues)
	values = values[1 : len(values)-1]
	return fmt.Sprintf("%-14s %s", c.Name, values)
}

func all(min, max int64) []int64 {
	if min > max {
		return []int64{}
	}
	vals := make([]int64, max-min+1)
	for i := min; i <= max; i++ {
		vals[i-min] = i
	}
	return vals
}

func parseRange(s string) (min int64, max int64, err error) {
	rangeLimits := strings.Split(s, "-")
	if len(rangeLimits) != 2 {
		return 0, 0, fmt.Errorf("wrong amount of range parameters")
	}
	min, err = strconv.ParseInt(rangeLimits[0], 10, 64)
	if err != nil {
		return 0, 0, errors.Wrap(err, "parsing min range value")
	}

	max, err = strconv.ParseInt(rangeLimits[1], 10, 64)
	if err != nil {
		return 0, 0, errors.Wrap(err, "parsing max range value")
	}
	return
}

func intifyBoolSieve(offset int64, sieve []bool) []int64 {
	vals := make([]int64, 0, len(sieve))

	for i, v := range sieve {
		if v {
			vals = append(vals, int64(i)+offset)
		}
	}
	return vals
}

func (c *CronValue) parse() error {
	timeSieve := make([]bool, c.Max-c.Min+1)
	for _, cronTimer := range strings.Split(c.Value, ",") {
		switch {
		case cronTimer == "*":
			for i := c.Min; i <= c.Max; i++ {
				c.parsedValues = all(c.Min, c.Max)
				return nil
			}
		case strings.HasPrefix(cronTimer, "*/"):
		case strings.Contains(cronTimer, "-"):
			min, max, err := parseRange(cronTimer)
			if err != nil {
				return err
			}
			if min > max {
				return fmt.Errorf("range in wrong order, is: %s, should be: %d-%d", c.Value, max, min)
			}
			if min < c.Min || max > c.Max {
				return fmt.Errorf("range out of cron value range: name: %s, value: %s, range: %d-%d", c.Name, c.Value, c.Min, c.Max)
			}
			for i := min; i <= max; i++ {
				timeSieve[i-c.Min] = true
			}
		default:

		}
	}
	c.parsedValues = intifyBoolSieve(c.Min, timeSieve)
	return nil
}

type Cron struct {
	Minute     CronValue
	Hour       CronValue
	DayOfMonth CronValue
	Month      CronValue
	DayOfWeek  CronValue
	Command    string
}

func (c *Cron) parse() {
	c.Minute.parse()
	c.Hour.parse()
	c.DayOfMonth.parse()
	c.Month.parse()
	c.DayOfWeek.parse()
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
			Name:  "minute",
			Value: args[0],
			Min:   0,
			Max:   59,
		},
		Hour: CronValue{
			Name:  "hour",
			Value: args[1],
			Min:   0,
			Max:   23,
		},
		DayOfMonth: CronValue{
			Name:  "day of month",
			Value: args[2],
			Min:   1,
			Max:   31,
		},
		Month: CronValue{
			Name:  "month",
			Value: args[3],
			Min:   1,
			Max:   12,
		},
		DayOfWeek: CronValue{
			Name:  "day of week",
			Value: args[4],
			Min:   0,
			Max:   6,
		},
		Command: args[5],
	}
	c.parse()
	return c, nil
}
