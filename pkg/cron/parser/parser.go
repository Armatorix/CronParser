package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errWrongFormat = errors.New("wrong format")
	errMinGTMax    = errors.New("min greater than max")
	errStepTooBig  = errors.New("step too big")
)

// ParseRange parses string from format "${min}-${max}" as min, max values
// where min, max are integers
// return error in case of wrong format or when min>max.
func ParseRange(s string) (min int64, max int64, err error) {
	rangeLimits := strings.Split(s, "-")
	if len(rangeLimits) != 2 {
		return 0, 0, errWrongFormat
	}
	min, err = strconv.ParseInt(rangeLimits[0], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("parse min value: %w", err)
	}

	max, err = strconv.ParseInt(rangeLimits[1], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("parse max value: %w", err)
	}

	if min > max {
		return min, max, fmt.Errorf("%w: min: %d, max: %d", errMinGTMax, min, max)
	}
	return min, max, nil
}

// ParseStep returnes values for min-max range
// with step parsed from s in format "*/${step}"
// where step is an integer
// return error in case of wrong format or step bigger than range.
func ParseStep(s string, min, max int64) ([]int64, error) {
	if !strings.HasPrefix(s, "*/") {
		return nil, errWrongFormat
	}
	if len(s) == 2 {
		return nil, fmt.Errorf("%w: missing step", errWrongFormat)
	}
	step, err := strconv.ParseInt(s[2:], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("%w: parse step value", err)
	}
	if max-min < step {
		return nil, errStepTooBig
	}
	vals := make([]int64, ((max-min)/step)+1)

	for i := range vals {
		vals[i] = min + int64(i)*step
	}
	return vals, nil
}
