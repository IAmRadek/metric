package money

import (
	"fmt"

	"github.com/IAmRadek/metric"
)

type Tax struct {
	metric.Quantity
}

func NewTax(percents float64, taxType TaxType) Tax {
	return Tax{
		Quantity: metric.NewQuantity(percents, taxType),
	}
}

func (t Tax) Rate() float64 {
	return t.Amount()
}

func (t Tax) Type() TaxType {
	return t.Metric().(TaxType)
}

func (t Tax) String() string {
	return fmt.Sprintf("%s (%.0f%%)", t.Type().Symbol(), t.Rate())
}

// Calculate returns how much tax there is
func (t Tax) Calculate(m Money) Money {
	q, _ := m.Multiply(t.Rate() / 100.0)

	return q
}
