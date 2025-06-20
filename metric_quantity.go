package metric

import (
	"fmt"
	"math/big"
)

type Quantity interface {
	Amount() float64
	Metric() Metric
	String() string

	// Add adds two Quantity objects
	// Precondition: both the target and the parameter Quantity objects must be in the same Metric
	// Returns a new Quantity object that has an amount equal to the sum of the amounts of the target Quantity object and the parameter Quantity object
	Add(Quantity) (Quantity, error)

	// Subtract subtracts one Quantity object from another
	// Precondition: both the target and the parameter Quantity objects must be in the same Metric
	// Returns a new Quantity object that has an amount equal to the amount of the target Quantity object minus the amount of the parameter Quantity object
	Subtract(Quantity) (Quantity, error)

	// Multiply multiplying the Quantity object by the multiplier
	// Returns a new Quantity object that has an amount equal to the amount of the target Quantity object multiplied by the multiplier
	Multiply(multiplier float64) (Quantity, error)

	// MultiplyBy multiplies two Quantity objects
	// Returns a new Quantity object that has an amount equal to the product of the amounts of the target Quantity object and the parameter Quantity object
	// The Metric of the returned Quantity object is a DerivedUnit given by the following equation:
	// T*P where T is the Unit of the target object and P is the Unit of the parameter object
	MultiplyBy(multiplier Quantity) (Quantity, error)

	// Round rounding the Quantity object given a RoundingPolicy
	// Returns a new Quantity object that has an amount equal to the amount of the target Quantity object rounded according to the RoundingPolicy
	Round(policy RoundingPolicy) (Quantity, error)

	// Divide dividing the Quantity object by the divisor
	// Returns a new Quantity object that has an amount equal to the amount of the target Quantity object divided by the divisor
	Divide(divisor float64) (Quantity, error)

	// DivideBy dividing the Quantity object by the divisor Quantity object
	// Divides one Quantity object by another Returns a new Quantity object that has an amount equal to the amount of the target Quantity object divided by the amount of the parameter Quantity object
	// The Metric of the returned Quantity object is a DerivedUnit given by the following equation: TP–1 where T is the Unit of the target object and P is the Unit of the parameter object
	DivideBy(divisor Quantity) (Quantity, error)

	// Equals compares two Quantity objects
	// Precondition: both the target and the parameter Quantity objects must be in the same Metric
	// Returns true if the amount of the target Quantity object is equal to the amount of the parameter Quantity object
	Equals(Quantity) (bool, error)

	// GreaterThan compares two Quantity objects
	// Precondition: both the target and the parameter Quantity objects must be in the same Metric
	// Returns true if the amount of the target Quantity object is greater than the amount of the parameter Quantity object
	GreaterThan(Quantity) (bool, error)

	// LessThan compares two Quantity objects
	// Precondition: both the target and the parameter Quantity objects must be in the same Metric
	// Returns true if the amount of the target Quantity object is less than the amount of the parameter Quantity object
	LessThan(Quantity) (bool, error)
}

type quantityImpl struct {
	amount float64
	metric Metric
}

func NewQuantity(amount float64, metric Metric) Quantity {
	return &quantityImpl{
		amount: amount,
		metric: metric,
	}
}

func (q *quantityImpl) Amount() float64 {
	return q.amount
}

func (q *quantityImpl) String() string {
	return fmt.Sprintf("%v %s", q.amount, q.metric)
}

func (q *quantityImpl) Metric() Metric {
	return q.metric
}

func (q *quantityImpl) Add(q2 Quantity) (Quantity, error) {
	if q.metric != q2.Metric() {
		return nil, ErrIncompatibleMetric{q.metric, q2.Metric()}
	}

	return NewQuantity(q.amount+q2.Amount(), q.metric), nil
}

func (q *quantityImpl) Subtract(q2 Quantity) (Quantity, error) {
	if q.metric != q2.Metric() {
		return nil, ErrIncompatibleMetric{q.metric, q2.Metric()}
	}

	return NewQuantity(q.amount-q2.Amount(), q.metric), nil
}

func (q *quantityImpl) Multiply(multiplier float64) (Quantity, error) {
	return NewQuantity(q.amount*multiplier, q.metric), nil
}

func (q *quantityImpl) MultiplyBy(q2 Quantity) (Quantity, error) {
	name := fmt.Sprintf("%s*%s", q.metric, q2.Metric())
	definition := fmt.Sprintf("Describes the product of %s and %s", q.metric, q2.Metric())
	symbol := fmt.Sprintf("%s*%s", q.metric.Symbol(), q2.Metric().Symbol())

	du := NewDerivedUnit(
		name,
		definition,
		symbol,
		nil,
		NewDerivedUnitTerm(q.metric, 1),
		NewDerivedUnitTerm(q2.Metric(), 1),
	)

	return NewQuantity(q.amount*q2.Amount(), du), nil
}

func (q *quantityImpl) Round(policy RoundingPolicy) (Quantity, error) {
	return NewQuantity(policy.Round(q.amount), q.metric), nil
}

func (q *quantityImpl) Divide(divisor float64) (Quantity, error) {
	if divisor == 0 {
		return nil, ErrDivisionByZero
	}
	return NewQuantity(q.amount/divisor, q.metric), nil
}

func (q *quantityImpl) DivideBy(divisor Quantity) (Quantity, error) {
	name := fmt.Sprintf("%s/%s", q.metric, divisor.Metric())
	definition := fmt.Sprintf("Describes the ratio between %s and %s", q.metric, divisor.Metric())
	symbol := fmt.Sprintf("%s/%s", q.metric.Symbol(), divisor.Metric().Symbol())

	du := NewDerivedUnit(
		name,
		definition,
		symbol,
		nil,
		NewDerivedUnitTerm(q.metric, 1),
		NewDerivedUnitTerm(divisor.Metric(), -1),
	)

	return NewQuantity(q.amount/divisor.Amount(), du), nil
}

func (q *quantityImpl) Equals(q2 Quantity) (bool, error) {
	if q.metric != q2.Metric() {
		return false, ErrIncompatibleMetric{q.metric, q2.Metric()}
	}

	result := big.NewFloat(q.amount).Cmp(big.NewFloat(q2.Amount()))
	return result == 0, nil
}

func (q *quantityImpl) GreaterThan(q2 Quantity) (bool, error) {
	if q.metric != q2.Metric() {
		return false, ErrIncompatibleMetric{q.metric, q2.Metric()}
	}

	result := big.NewFloat(q.amount).Cmp(big.NewFloat(q2.Amount()))
	return result > 0, nil
}

func (q *quantityImpl) LessThan(q2 Quantity) (bool, error) {
	if q.metric != q2.Metric() {
		return false, ErrIncompatibleMetric{q.metric, q2.Metric()}
	}

	result := big.NewFloat(q.amount).Cmp(big.NewFloat(q2.Amount()))
	return result < 0, nil
}
