package metric_test

import (
	"testing"

	"github.com/IAmRadek/metric"
	isser "github.com/matryer/is"
)

func TestUnitConverter(t *testing.T) {
	is := isser.New(t)

	celsius := metric.NewQuantity(1, metric.Celsius)

	kelvins, err := metric.UnitConverter.Convert(celsius, metric.Kelvin)
	is.NoErr(err)

	expKelvins := metric.NewQuantity(274.15, metric.Kelvin)
	equals, err := kelvins.Equals(expKelvins)
	is.NoErr(err)
	is.True(equals)
}

func TestDirectUnitConverter(t *testing.T) {
	is := isser.New(t)

	celsius := metric.NewQuantity(1, metric.Celsius)

	kelvin, err := metric.CelsiusToKelvin.Convert(celsius)
	is.NoErr(err)

	is.True(kelvin.Metric() == metric.Kelvin)

	c2, err := metric.KelvinToCelsius.Convert(kelvin)
	is.NoErr(err)

	is.True(c2.Metric() == metric.Celsius)
	is.True(c2.Amount() == celsius.Amount())
}
