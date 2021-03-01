package cron

import (
	"fmt"
	"os"
	"strings"
)

// CronValue provides minimal entity from cron object
// it handles minute, hour, day, month day of month
type CronValue struct {
	Name  string
	Value string
	Min   int
	Max   int

	parsedValues []int
}

func (c CronValue) String() string {
	values := fmt.Sprint(c.parsedValues)
	values = values[1 : len(values)-1]
	return fmt.Sprintf("%-14s %s", c.Name, values)
}

func (c *CronValue) validate() error {
	return nil
}

func all(min, max int) []int {
	if min > max {
		return []int{}
	}
	vals := make([]int, max-min+1)
	for i := min; i <= max; i++ {
		vals[i-min] = i
	}
	return vals
}

func (c *CronValue) parse() {
	// timeSieve := make([]bool, c.Max-c.Min+1)
	// times := make([]int, c.Max-c.Min+1)
	for _, cronTimer := range strings.Split(c.Value, ",") {
		switch {
		case cronTimer == "*":
			for i := c.Min; i <= c.Max; i++ {
				c.parsedValues = all(c.Min, c.Max)
				return
			}
		}
	}
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
