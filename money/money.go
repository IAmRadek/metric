package money

import (
	"fmt"

	"github.com/IAmRadek/metric"
	"github.com/govalues/decimal"
)

// Money is a special monetary Quantity.
type Money struct {
	amount   decimal.Decimal
	currency Currency
}

func NewMoney(minorUnit int64, currency Currency) Money {
	return Money{
		amount:   decimal.MustNew(minorUnit, currency.Decimal()),
		currency: currency,
	}
}

// MinorUnit returns the amount in the smallest subdivision of the Currency.
func (m Money) MinorUnit() uint64 {
	return m.amount.Coef()
}

// String returns a formatted string representation of the Money object in the format "{amount} {currency code}"
func (m Money) String() string {
	w, f, _ := m.amount.Int64(m.currency.Decimal())

	return fmt.Sprintf("%d.%d %s", w, f, m.Currency().Code())
}

// Metric returns the Currency associated with the Money object.
func (m Money) Metric() metric.Metric {
	return m.currency
}

// Currency returns the Currency of the Money object.
func (m Money) Currency() Currency {
	return m.currency
}

// AfterTax returns Money with added Tax value.
func (m Money) AfterTax(t Tax) Money {
	taxes := t.Calculate(m)
	afterTax, _ := m.Add(taxes)

	return afterTax
}

// AfterTaxes Returns Money with added Tax values.
func (m Money) AfterTaxes(taxes []Tax) Money {
	var out = m
	for _, t := range taxes {
		out, _ = m.Add(t.Calculate(m))
	}

	return out
}

// Add adds two Money objects
// Precondition: both the target and the parameter Money objects must be in the same Currency
// Returns a new Money object that has an amount equal to the sum of the amounts of the target Money object and the parameter Money object
func (m Money) Add(m2 Money) (Money, error) {
	if m.currency != m2.Currency() {
		return Money{}, metric.ErrIncompatibleMetric{M1: m.Metric(), M2: m2.Metric()}
	}

	cents, err := m.amount.Add(m2.amount)
	if err != nil {
		return Money{}, fmt.Errorf("error adding Money objects: %w", err)
	}
	return Money{cents, m.currency}, nil
}

// Subtract subtracts one Money object from another
// Precondition: both the target and the parameter Money objects must be in the same Currency
// Returns a new Money object that has an amount equal to the amount of the target Money object minus the amount of the parameter Money object
func (m Money) Subtract(m2 Money) (Money, error) {
	if m.currency != m2.Currency() {
		return Money{}, metric.ErrIncompatibleMetric{M1: m.Metric(), M2: m2.Metric()}
	}

	cents, err := m.amount.Sub(m2.amount)
	if err != nil {
		return Money{}, fmt.Errorf("error subtracting Money objects: %w", err)
	}

	return Money{cents, m.currency}, nil
}

// Multiply multiplying the Money object by the multiplier
// Returns a new Money object that has an amount equal to the amount of the target Money object multiplied by the multiplier
func (m Money) Multiply(multiplier float64) (Money, error) {
	mul, _ := decimal.NewFromFloat64(multiplier)

	cents, err := m.amount.MulExact(mul, m.currency.Decimal())
	if err != nil {
		return Money{}, fmt.Errorf("error multiplying Money objects: %w", err)
	}

	return Money{cents, m.currency}, nil
}

// Divide dividing the Money object by the divisor
// Returns a new Money object that has an amount equal to the amount of the target Money object divided by the divisor
func (m Money) Divide(divisor float64) (Money, error) {
	div, err := decimal.NewFromFloat64(divisor)
	if err != nil {
		return Money{}, fmt.Errorf("invalid divisor: %w", err)
	}

	am, err := m.amount.Quo(div)
	if err != nil {
		return Money{}, fmt.Errorf("dividing: %w", err)
	}

	return Money{
		amount:   am,
		currency: m.currency,
	}, nil
}

// DivideBy dividing the Money object by the divisor metric.Quantity object
// Returns a new metric.Quantity object that has an amount equal to the amount of the target Money object divided by the amount of the parameter metric.Quantity object
// The Metric of the returned metric.Quantity object is a metric.DerivedUnit given by the following equation: TP–1 where T is the metric.Unit of the target object and P is the metric.Unit of the parameter object
func (m Money) DivideBy(divisor metric.Quantity) (metric.Quantity, error) {
	name := fmt.Sprintf("%s/%s", m.currency, divisor.Metric())
	definition := fmt.Sprintf("Describes the cost per unit, %s per %s", m.currency, divisor.Metric())
	symbol := fmt.Sprintf("%s/%s", m.currency.Symbol(), divisor.Metric().Symbol())

	du := metric.NewDerivedUnit(
		name,
		definition,
		symbol,
		nil,
		metric.NewDerivedUnitTerm(m.currency, 1),
		metric.NewDerivedUnitTerm(divisor.Metric(), -1),
	)

	div, err := decimal.NewFromFloat64(divisor.Amount())
	if err != nil {
		return nil, fmt.Errorf("invalid divisor: %w", err)
	}

	am, err := m.amount.Quo(div)
	if err != nil {
		return nil, fmt.Errorf("dividing: %w", err)
	}

	f, ok := am.Float64()
	if !ok {
		return nil, fmt.Errorf("error converting to float64: %w", err)
	}

	return metric.NewQuantity(f, du), nil
}

// Equals compares two Money objects
// Precondition: both the target and the parameter Money objects must be in the same Currency
// Returns true if the amount of the target Money object is equal to the amount of the parameter Money object
func (m Money) Equals(m2 Money) (bool, error) {
	if m.currency != m2.Currency() {
		return false, metric.ErrIncompatibleMetric{M1: m.Metric(), M2: m2.Metric()}
	}

	return m.amount.Equal(m2.amount), nil
}

// GreaterThan compares two Money objects
// Precondition: both the target and the parameter Money objects must be in the same Currency
// Returns true if the amount of the target Money object is greater than the amount of the parameter Money object
func (m Money) GreaterThan(m2 Money) (bool, error) {
	if m.currency != m2.Currency() {
		return false, metric.ErrIncompatibleMetric{M1: m.Metric(), M2: m2.Metric()}
	}

	return !m.amount.Less(m2.amount), nil
}

// LessThan compares two Money objects
// Precondition: both the target and the parameter Money objects must be in the same Currency
// Returns true if the amount of the target Money object is less than the amount of the parameter Money object
func (m Money) LessThan(m2 Money) (bool, error) {
	if m.currency != m2.Currency() {
		return false, metric.ErrIncompatibleMetric{M1: m.Metric(), M2: m2.Metric()}
	}

	return m.amount.Less(m2.amount), nil
}

// IsZero returns true if the Money amount is zero
func (m Money) IsZero() bool {
	return m.amount.IsZero()
}
