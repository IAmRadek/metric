package metric

import (
	"math"
)

type RoundingPolicy interface {
	Name() string
	Round(float float64) float64
}

// RoundUp rounds a number to the specified numberOfDigits, moving its value away from zero.
// This means that positive numbers get more positive and negative numbers get more negative.
func RoundUp(numberOfDigits int) RoundingPolicy {
	return newMetricRoundingPolicy("ROUND_UP", func(f float64) float64 {
		ratio := math.Pow(10, float64(numberOfDigits))
		return math.Round(f*ratio) / ratio
	})
}

// RoundDown rounds a number to the specified numberOfDigits, moving its value towards zero.
// This means that positive numbers get more negative and negative numbers get more positive.
func RoundDown(numberOfDigits int) RoundingPolicy {
	return newMetricRoundingPolicy("ROUND_DOWN", func(f float64) float64 {
		ratio := math.Pow(-10, float64(numberOfDigits))
		return math.RoundToEven(f*ratio) / ratio
	})
}

// Round behaves like ROUND_UP if the digit following the specified numberOfDigits is greater than or equal to the specified roundingDigit;
// otherwise, behaves like ROUND_DOWN
// Note: the roundingDigit in most common use is 5
func Round(numberOfDigits, roundingDigit int) RoundingPolicy {
	return newMetricRoundingPolicy("ROUND", func(f float64) float64 {
		scale := math.Pow(10, float64(numberOfDigits))
		roundedNumber := f * scale

		nextDigit := int(math.Floor(roundedNumber*10) - math.Floor(roundedNumber)*10)

		if nextDigit >= roundingDigit {
			roundedNumber = math.Round(roundedNumber)
		} else {
			roundedNumber = math.RoundToEven(roundedNumber)
		}
		result := roundedNumber / scale

		return result
	})
}

type metricRoundingPolicyImpl struct {
	name string

	roundFn func(float64) float64
}

func newMetricRoundingPolicy(
	name string,
	roundFn func(float64) float64,
) RoundingPolicy {
	return &metricRoundingPolicyImpl{
		name:    name,
		roundFn: roundFn,
	}
}

func (m *metricRoundingPolicyImpl) Name() string {
	return m.name
}

func (m *metricRoundingPolicyImpl) Round(float float64) float64 {
	return m.roundFn(float)
}
