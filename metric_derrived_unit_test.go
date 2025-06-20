package metric_test

import (
	"testing"

	"github.com/IAmRadek/metric"
	isser "github.com/matryer/is"
)

func TestDerivedUnit(t *testing.T) {
	is := isser.New(t)

	// Create a system of units for testing
	system := metric.NewSystemOfUnits("Test System", "Test Body")

	// Create a derived unit
	derivedUnit := metric.NewDerivedUnit(
		"Test Derived Unit",
		"Test Definition",
		"TDU",
		system,
		metric.NewDerivedUnitTerm(metric.Meter, 1),
		metric.NewDerivedUnitTerm(metric.Second, -1),
	)

	// Test methods
	is.Equal(derivedUnit.Name(), "Test Derived Unit")
	is.Equal(derivedUnit.Definition(), "Test Definition")
	is.Equal(derivedUnit.Symbol(), "TDU")
	is.Equal(derivedUnit.String(), "TDU")
	is.Equal(derivedUnit.SystemOfUnits(), system)

	// Test Terms method
	terms := derivedUnit.Terms()
	is.Equal(len(terms), 2)
	is.Equal(terms[0].Metric(), metric.Meter)
	is.Equal(terms[0].Exponent(), 1)
	is.Equal(terms[1].Metric(), metric.Second)
	is.Equal(terms[1].Exponent(), -1)
}

func TestDerivedUnitTerm(t *testing.T) {
	is := isser.New(t)

	// Create a derived unit term
	term := metric.NewDerivedUnitTerm(metric.Meter, 2)

	// Test methods
	is.Equal(term.Name(), metric.Meter.Name())
	is.Equal(term.Definition(), metric.Meter.Definition())
	is.Equal(term.Symbol(), "m^2")
	is.Equal(term.String(), "m^2")
	is.Equal(term.Exponent(), 2)
	is.Equal(term.Metric(), metric.Meter)
}