package cron

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Armatorix/CronParser/pkg/cron/parser"
	"github.com/Armatorix/CronParser/pkg/existencemap"
)

// CronValue provides minimal entity from cron object
// it handles minute, hour, day, month day of month.
type CronValue struct {
	name  string
	value string
	min   int64
	max   int64

	parsedValues []int64
}

func NewCronValue(name, value string, min, max int64) (*CronValue, error) {
	cv := &CronValue{
		name:  name,
		value: value,
		min:   min,
		max:   max,
	}
	err := cv.parse()
	if err != nil {
		return nil, fmt.Errorf("%s parsing failed: %w", cv.name, err)
	}

	return cv, nil
}

func (c CronValue) String() string {
	values := fmt.Sprint(c.parsedValues)
	values = values[1 : len(values)-1]
	return fmt.Sprintf("%-14s %s", c.name, values)
}

func (c *CronValue) parse() error {
	existence, err := existencemap.New(c.min, c.max)
	if err != nil {
		return err
	}
	for _, cronTimer := range strings.Split(c.value, ",") {
		switch {
		case cronTimer == "*":
			for i := c.min; i <= c.max; i++ {
				existence.AllExists()
				c.parsedValues = existence.ToInt64Slice()
				return nil
			}
		case strings.HasPrefix(cronTimer, "*/"):
			vals, err := parser.ParseStep(cronTimer, c.min, c.max)
			if err != nil {
				return err
			}
			if err = existence.ApplySlice(vals); err != nil {
				return err
			}

		case strings.Contains(cronTimer, "-"):
			min, max, err := parser.ParseRange(cronTimer)
			if err != nil {
				return err
			}

			if err = existence.ApplyRange(min, max); err != nil {
				return err
			}
		default:
			v, err := strconv.ParseInt(cronTimer, 10, 64)
			if err != nil {
				return fmt.Errorf("single value parse failed: %w", err)
			}

			if err = existence.ApplyNumber(v); err != nil {
				return err
			}
		}
	}
	c.parsedValues = existence.ToInt64Slice()
	return nil
}
