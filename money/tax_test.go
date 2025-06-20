package money_test

import (
	"testing"

	"github.com/IAmRadek/metric/money"
	isser "github.com/matryer/is"
)

func TestTax(t *testing.T) {
	tests := []struct {
		name  string
		tax   money.Tax
		check func(is *isser.I, tax money.Tax)
	}{
		{
			name: "Type",
			tax:  money.NewTax(23, money.VAT),
			check: func(is *isser.I, tax money.Tax) {
				is.Equal(tax.Type(), money.VAT)
				is.Equal(tax.Type().Type(), "VAT")
			},
		},
		{
			name: "String",
			tax:  money.NewTax(23, money.VAT),
			check: func(is *isser.I, tax money.Tax) {
				is.Equal(tax.String(), "VAT (23%)")
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			is := isser.New(t)
			tt.check(is, tt.tax)
		})
	}
}

func TestTaxType(t *testing.T) {
	tests := []struct {
		name    string
		taxType money.TaxType
		check   func(is *isser.I, taxType money.TaxType)
	}{
		{
			name:    "Type",
			taxType: money.VAT,
			check: func(is *isser.I, taxType money.TaxType) {
				is.Equal(taxType.Type(), "VAT")
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			is := isser.New(t)
			tt.check(is, tt.taxType)
		})
	}
}
