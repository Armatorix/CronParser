package cron

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Armatorix/CronParser/pkg/cron/parser"
	"github.com/Armatorix/CronParser/pkg/existancemap"
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

func (c *CronValue) parse() error {
	existance, err := existancemap.New(c.Min, c.Max)
	if err != nil {
		return err
	}
	for _, cronTimer := range strings.Split(c.Value, ",") {
		switch {
		case cronTimer == "*":
			for i := c.Min; i <= c.Max; i++ {
				existance.AllExists()
				c.parsedValues = existance.ToInt64Slice()
				return nil
			}
		case strings.HasPrefix(cronTimer, "*/"):
			vals, err := parser.ParseStep(cronTimer, c.Min, c.Max)
			if err != nil {
				return err
			}
			if err = existance.ApplySlice(vals); err != nil {
				return err
			}

		case strings.Contains(cronTimer, "-"):
			min, max, err := parser.ParseRange(cronTimer)
			if err != nil {
				return err
			}

			if err = existance.ApplyRange(min, max); err != nil {
				return err
			}
		default:
			v, err := strconv.ParseInt(cronTimer, 10, 64)
			if err != nil {
				return errors.Wrap(err, "single value parse failed")
			}

			if err = existance.ApplyNumber(v); err != nil {
				return err
			}
		}
	}
	c.parsedValues = existance.ToInt64Slice()
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

// For now return first occured error, can be extended to show all of them
func (c *Cron) parse() error {
	cronEntities := []CronValue{
		c.Minute,
		c.Hour,
		c.DayOfMonth,
		c.Month,
		c.DayOfWeek,
	}
	for _, cronEntity := range cronEntities {
		if err := cronEntity.parse(); err != nil {
			return errors.WithMessagef(err, "%s parsing failed", cronEntity.Name)
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
	err := c.parse()
	if err != nil {
		return nil, err
	}
	return c, nil
}
