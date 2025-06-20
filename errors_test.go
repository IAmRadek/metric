package metric_test

import (
	"testing"

	"github.com/IAmRadek/metric"
	isser "github.com/matryer/is"
)

func TestErrIncompatibleMetric_Error(t *testing.T) {
	is := isser.New(t)

	// Create a mock metric for testing
	metric1 := mockMetric{symbol: "m"}
	metric2 := mockMetric{symbol: "kg"}

	// Create the error
	err := metric.ErrIncompatibleMetric{M1: metric1, M2: metric2}

	// Test the Error method
	is.Equal(err.Error(), `metric "m" is not compatible with "kg"`)
}

// mockMetric is a simple implementation of the Metric interface for testing
type mockMetric struct {
	symbol string
}

func (m mockMetric) Name() string {
	return "Mock Metric"
}

func (m mockMetric) Definition() string {
	return "Mock Definition"
}

func (m mockMetric) Symbol() string {
	return m.symbol
}

func (m mockMetric) String() string {
	return m.symbol
}