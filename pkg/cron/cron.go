package cron

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	errIncorrectCmdArgsLen    = errors.New("incorrect number of command line argument")
	errIncorrectCmdCronArgLen = errors.New("incorrect number of cron arguments")
)

type Cron struct {
	Minute     *CronValue
	Hour       *CronValue
	DayOfMonth *CronValue
	Month      *CronValue
	DayOfWeek  *CronValue
	Command    string
}

func (c Cron) String() string {
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%-14s %s\n",
		c.Minute, c.Hour, c.DayOfMonth, c.Month, c.DayOfWeek, "command", c.Command)
}
func New(args []string) (*Cron, error) {
	if len(args) != 6 {
		return nil, fmt.Errorf("%w: length: %d", errIncorrectCmdCronArgLen, len(args))
	}

	c := &Cron{
		Command: args[5],
	}
	var err error
	if c.Minute, err = NewCronValue("minute", args[0], 0, 59); err != nil {
		return nil, err
	}
	if c.Hour, err = NewCronValue("hour", args[1], 0, 23); err != nil {
		return nil, err
	}
	if c.DayOfMonth, err = NewCronValue("day of month", args[2], 1, 31); err != nil {
		return nil, err
	}
	if c.Month, err = NewCronValue("month", args[3], 1, 12); err != nil {
		return nil, err
	}
	if c.DayOfWeek, err = NewCronValue("day of week", args[4], 0, 6); err != nil {
		return nil, err
	}

	return c, nil
}
func NewFromOsArgs() (*Cron, error) {
	osArgs := os.Args
	if len(osArgs) != 2 {
		return nil, fmt.Errorf("%w: length %d", errIncorrectCmdArgsLen, len(osArgs))
	}
	args := strings.Split(osArgs[1], " ")

	c, err := New(args)
	if err != nil {
		return nil, err
	}
	return c, nil
}
