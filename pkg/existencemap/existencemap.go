package existencemap

import (
	"github.com/pkg/errors"
)

var (
	errMinGTMax   = errors.New("min greater than max")
	errOutOfBound = errors.New("out of bound")
)

type ExistenceMap struct {
	existence []bool
	min       int64
	max       int64
}

// New creates existenceMap with min-max range handling
// returns error in case of max<min.
func New(min, max int64) (*ExistenceMap, error) {
	if min > max {
		return nil, errMinGTMax
	}
	return &ExistenceMap{
		min:       min,
		max:       max,
		existence: make([]bool, max-min+1),
	}, nil
}

// ApplyNumber marks value as existing
// returns error if value is out of bound.
func (e *ExistenceMap) ApplyNumber(v int64) error {
	if v < e.min || v > e.max {
		return errors.WithMessagef(errOutOfBound, "min: %d, max: %d, value: %d", e.min, e.max, v)
	}
	e.existence[int(v-e.min)] = true
	return nil
}

// ApplySlice marks all values from slice as existing
// returns error if any value is out of bound.
func (e *ExistenceMap) ApplySlice(vals []int64) error {
	for _, v := range vals {
		if err := e.ApplyNumber(v); err != nil {
			return err
		}
	}
	return nil
}

// ApplyRange marks all values from slice as existing
// returns error if any value is out of bound.
func (e *ExistenceMap) ApplyRange(min, max int64) error {
	if min < e.min || max > e.max {
		return errors.WithMessagef(errOutOfBound, "existence %d-%d, applied %d-%d", e.min, e.max, min, max)
	}
	for v := min; v <= max; v++ {
		e.existence[v-e.min] = true
	}
	return nil
}
func (e *ExistenceMap) AllExists() {
	for i := range e.existence {
		e.existence[i] = true
	}
}

func (e ExistenceMap) ToInt64Slice() []int64 {
	s := make([]int64, 0, int(e.max-e.min+1))
	for i, v := range e.existence {
		if v {
			s = append(s, int64(i)+e.min)
		}
	}
	return s
}
