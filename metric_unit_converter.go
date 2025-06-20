package metric

import (
	"errors"
	"fmt"
)

var (
	ErrNoConversion    = errors.New("no conversion found")
	ErrMetricIsNotUnit = errors.New("metric is not a unit")
)

var (
	UnitConverter = &defaultUnitConverter{
		conversions: make(map[Unit]map[Unit]StandardConversion),
	}
)

type defaultUnitConverter struct {
	conversions map[Unit]map[Unit]StandardConversion
}

var (
	// CelsiusToKelvin represents a conversion from Celsius to Kelvin adjusting for absolute zero (-273.15°C).
	CelsiusToKelvin = NewStandardConversion(Celsius, Kelvin, func(quantity Quantity) (Quantity, error) {
		const absoluteZero = -273.15
		return NewQuantity(quantity.Amount()-absoluteZero, Kelvin), nil
	})

	// KelvinToCelsius represents a conversion from Kelvin to Celsius accounting for absolute zero (-273.15°C).
	KelvinToCelsius = NewStandardConversion(Kelvin, Celsius, func(quantity Quantity) (Quantity, error) {
		const absoluteZero = -273.15
		return NewQuantity(quantity.Amount()+absoluteZero, Celsius), nil
	})
)

func (c *defaultUnitConverter) Convert(quantity Quantity, target Unit) (Quantity, error) {
	unit, ok := quantity.Metric().(Unit)
	if !ok {
		return nil, fmt.Errorf("%w: metric %q is not a unit", ErrMetricIsNotUnit, quantity.Metric().Name())
	}

	conversion, err := c.getConversion(unit, target)
	if err != nil {
		return nil, err
	}

	return conversion.Convert(quantity)
}

func (c *defaultUnitConverter) getConversion(sourceUnit, targetUnit Unit) (StandardConversion, error) {
	conversion, ok := c.conversions[targetUnit][sourceUnit]
	if !ok {
		return StandardConversion{}, ErrNoConversion
	}

	return conversion, nil
}

type StandardConversion struct {
	conversionFn func(Quantity) (Quantity, error)

	sourceUnit Unit
	targetUnit Unit
}

// NewStandardConversion creates a new StandardConversion instance.
// It registers the conversion in the default UnitConverter.
func NewStandardConversion(sourceUnit, targetUnit Unit, conversionFn func(Quantity) (Quantity, error)) StandardConversion {
	sc := StandardConversion{
		conversionFn: conversionFn,
		sourceUnit:   sourceUnit,
		targetUnit:   targetUnit,
	}

	if _, ok := UnitConverter.conversions[targetUnit]; !ok {
		UnitConverter.conversions[targetUnit] = make(map[Unit]StandardConversion)
	}

	UnitConverter.conversions[targetUnit][sourceUnit] = sc

	return sc
}

func (c StandardConversion) Convert(source Quantity) (Quantity, error) {
	if source.Metric() != c.sourceUnit {
		return nil, ErrMetricIsNotUnit
	}

	convertedAmount, err := c.conversionFn(source)
	if err != nil {
		return nil, err
	}

	return NewQuantity(convertedAmount.Amount(), c.targetUnit), nil
}
