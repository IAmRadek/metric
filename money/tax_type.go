package money

import (
	"github.com/IAmRadek/metric"
)

type TaxType interface {
	metric.Metric
	Type() string
}

type taxTypeImpl struct {
	metric.Metric
}

func (tax taxTypeImpl) MarshalText() (text []byte, err error) {
	return []byte(tax.Symbol()), nil
}

var TaxTypes = make(map[string]TaxType)

func NewTaxType(name, definition, symbol string) TaxType {
	tax := taxTypeImpl{
		Metric: metric.NewMetric(name, definition, symbol),
	}
	TaxTypes[name] = tax

	return tax
}

var (
	VAT = NewTaxType("VAT", "Value-added tax", "VAT")
)

func (tax taxTypeImpl) Type() string {
	return tax.Symbol()
}
