package metric_test

import (
	"testing"

	"github.com/IAmRadek/metric"
	isser "github.com/matryer/is"
)

func TestMetric(t *testing.T) {
	is := isser.New(t)

	// Test NewMetric and all methods of metricImpl
	m := metric.NewMetric("Length", "The measure of distance", "m")
	
	is.Equal(m.Name(), "Length")
	is.Equal(m.Definition(), "The measure of distance")
	is.Equal(m.Symbol(), "m")
	is.Equal(m.String(), "m") // String() returns Symbol()
}

func TestSystemOfUnits(t *testing.T) {
	is := isser.New(t)

	// Test Name and StandardizationBody methods
	s := metric.NewSystemOfUnits("Imperial", "British Empire")
	
	is.Equal(s.Name(), "Imperial")
	is.Equal(s.StandardizationBody(), "British Empire")
}