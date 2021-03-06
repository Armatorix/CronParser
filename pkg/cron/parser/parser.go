package parser

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var (
	errWrongFormat = errors.New("wrong format")
	errMinGTMax    = errors.New("min greater than max")
)

// ParseRange parses string from format "%d-%d" as min, max values
// return error in case of wrong format or when min>max
func ParseRange(s string) (min int64, max int64, err error) {
	rangeLimits := strings.Split(s, "-")
	if len(rangeLimits) != 2 {
		return 0, 0, errWrongFormat
	}
	min, err = strconv.ParseInt(rangeLimits[0], 10, 64)
	if err != nil {
		return 0, 0, errors.WithMessage(err, "parse min value")
	}

	max, err = strconv.ParseInt(rangeLimits[1], 10, 64)
	if err != nil {
		return 0, 0, errors.WithMessage(err, "parse min value")
	}

	if min > max {
		return min, max, errors.WithMessagef(errMinGTMax, "min: %d, max: %d", min, max)
	}
	return min, max, nil
}
