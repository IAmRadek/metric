package metric_test

import (
	"testing"

	"github.com/IAmRadek/metric"
	isser "github.com/matryer/is"
)

func TestSiBaseUnit_SystemOfUnits(t *testing.T) {
	is := isser.New(t)

	units := metric.SISystemOfUnits.Units()

	is.Equal(len(units), 16)

	is.True(containsUnit(units, metric.Meter))
	is.True(containsUnit(units, metric.Kilogram))
	is.True(containsUnit(units, metric.Second))
	is.True(containsUnit(units, metric.Ampere))
	is.True(containsUnit(units, metric.Kelvin))
	is.True(containsUnit(units, metric.Mole))
	is.True(containsUnit(units, metric.Candela))
	is.True(containsUnit(units, metric.Area))
	is.True(containsUnit(units, metric.Volume))
	is.True(containsUnit(units, metric.Speed))
	is.True(containsUnit(units, metric.Celsius))
	is.True(containsUnit(units, metric.Lux))
	is.True(containsUnit(units, metric.Lumen))
	is.True(containsUnit(units, metric.Steradian))
	is.True(containsUnit(units, metric.Radian))
	is.True(containsUnit(units, metric.Watt))
}

func TestSIBaseUnit_Methods(t *testing.T) {
	is := isser.New(t)

	// Test methods of Meter (an SIBaseUnit)
	is.Equal(metric.Meter.Name(), "meter")
	is.Equal(metric.Meter.Definition(), "The meter is the length of the path travelled by light in vacuum during a time interval of 1/299792458 of a second")
	is.Equal(metric.Meter.Symbol(), "m")
	is.Equal(metric.Meter.String(), "m")
	is.Equal(metric.Meter.SystemOfUnits(), metric.SISystemOfUnits)

	// Test methods of Kilogram (another SIBaseUnit)
	is.Equal(metric.Kilogram.Name(), "kilogram")
	is.Equal(metric.Kilogram.Definition(), "The kilogram is the unit of mass; it is equal to the mass of the international prototype of the kilogram")
	is.Equal(metric.Kilogram.Symbol(), "kg")
	is.Equal(metric.Kilogram.String(), "kg")
	is.Equal(metric.Kilogram.SystemOfUnits(), metric.SISystemOfUnits)
}

func TestSISystemOfUnits_Methods(t *testing.T) {
	is := isser.New(t)

	// Test methods of SISystemOfUnits
	is.Equal(metric.SISystemOfUnits.Name(), "SI")
	is.Equal(metric.SISystemOfUnits.StandardizationBody(), "BIPM")
}

func containsUnit(units []metric.Unit, unit metric.Unit) bool {
	for _, u := range units {
		if u == unit {
			return true
		}
	}
	return false
}
