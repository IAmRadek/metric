package metric

import (
	"errors"
	"fmt"
)

var (
	ErrDivisionByZero = errors.New("division by zero")
)

type ErrIncompatibleMetric struct {
	M1, M2 Metric
}

func (e ErrIncompatibleMetric) Error() string {
	return fmt.Sprintf("metric %q is not compatible with %q", e.M1, e.M2)
}
