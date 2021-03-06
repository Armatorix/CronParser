package existancemap

import (
	"github.com/pkg/errors"
)

var (
	errMinGTMax   = errors.New("min greater than max")
	errOutOfBound = errors.New("out of bound")
)

type ExistanceMap struct {
	existance []bool
	min       int64
	max       int64
}

// New creates ExistanceMap with min-max range handling
// returns error in case of max<min
func New(min, max int64) (*ExistanceMap, error) {

	if min > max {
		return nil, errMinGTMax
	}
	return &ExistanceMap{
		min:       min,
		max:       max,
		existance: make([]bool, max-min+1),
	}, nil
}

// ApplyNumber marks value as exising
// returns error if value is out of bound
func (e *ExistanceMap) ApplyNumber(v int64) error {
	if v < e.min || v > e.max {
		return errors.WithMessagef(errOutOfBound, "min: %d, max: %d, value: %d", e.min, e.max, v)
	}
	e.existance[int(v-e.min)] = true
	return nil
}

// ApplySlice marks all values from slice as exising
// returns error if any value is out of bound
func (e *ExistanceMap) ApplySlice(vals []int64) error {
	for _, v := range vals {
		if err := e.ApplyNumber(v); err != nil {
			return err
		}
	}
	return nil
}

// ApplyRange marks all values from slice as exising
// returns error if any value is out of bound
func (e *ExistanceMap) ApplyRange(min, max int64) error {
	if min < e.min || max > e.max {
		return errors.WithMessagef(errOutOfBound, "existance %d-%d, applied %d-%d", e.min, e.max, min, max)
	}
	for v := min; v <= max; v++ {
		e.existance[v-e.min] = true
	}
	return nil
}
func (e *ExistanceMap) AllExists() {
	for i := range e.existance {
		e.existance[i] = true
	}
}

func (e ExistanceMap) ToInt64Slice() []int64 {
	s := make([]int64, 0, int(e.max-e.min+1))
	for i, v := range e.existance {
		if v {
			s = append(s, int64(i)+e.min)
		}
	}
	return s
}
