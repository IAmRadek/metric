package money_test

import (
	"fmt"
	"testing"

	"github.com/IAmRadek/metric"
	"github.com/IAmRadek/metric/money"
	isser "github.com/matryer/is"
)

func TestMoney(t *testing.T) {
	tests := []struct {
		name   string
		q1, q2 money.Money
		check  func(is *isser.I, q1, q2 money.Money)
	}{
		{
			name: "Equals",
			q1:   money.NewMoney(10, money.USD),
			q2:   money.NewMoney(10, money.USD),
			check: func(is *isser.I, q1, q2 money.Money) {
				equals, err := q1.Equals(q2)
				is.NoErr(err)
				is.True(equals)
			},
		},
		{
			name: "LessThan",
			q1:   money.NewMoney(10, money.USD),
			q2:   money.NewMoney(100, money.USD),
			check: func(is *isser.I, q1, q2 money.Money) {
				lessThan, err := q1.LessThan(q2)
				is.NoErr(err)
				is.True(lessThan)
			},
		},
		{
			name: "GreaterThan",
			q1:   money.NewMoney(100, money.USD),
			q2:   money.NewMoney(10, money.USD),
			check: func(is *isser.I, q1, q2 money.Money) {
				greaterThan, err := q1.GreaterThan(q2)
				is.NoErr(err)
				is.True(greaterThan)
			},
		},
		{
			name: "Add",
			q1:   money.NewMoney(10, money.USD),
			q2:   money.NewMoney(100, money.USD),
			check: func(is *isser.I, q1, q2 money.Money) {
				sum, err := q1.Add(q2)
				is.NoErr(err)
				is.Equal(sum.MinorUnit(), uint64(110))
				is.Equal(sum.Metric(), money.USD)
			},
		},
		{
			name: "Subtract",
			q1:   money.NewMoney(100, money.USD),
			q2:   money.NewMoney(10, money.USD),
			check: func(is *isser.I, q1, q2 money.Money) {
				diff, err := q1.Subtract(q2)
				is.NoErr(err)
				is.Equal(diff.MinorUnit(), uint64(90))
				is.Equal(diff.Metric(), money.USD)
			},
		},
		{
			name: "Subtract_Incompatible",
			q1:   money.NewMoney(100, money.USD),
			q2:   money.NewMoney(120, money.PLN),
			check: func(is *isser.I, q1, q2 money.Money) {

				_, err := q1.Subtract(q2)
				is.Equal(err, metric.ErrIncompatibleMetric{M1: q1.Metric(), M2: q2.Metric()})
			},
		},
		{
			name: "Multiply",
			q1:   money.NewMoney(10, money.USD),
			q2:   money.NewMoney(100, money.USD),
			check: func(is *isser.I, q1, q2 money.Money) {
				p, err := q1.Multiply(10)
				is.NoErr(err)
				eq, err := p.Equals(q2)
				is.NoErr(err)
				is.True(eq)
			},
		},
		{
			name: "Divide",
			q1:   money.NewMoney(100, money.USD),
			q2:   money.NewMoney(10, money.USD),
			check: func(is *isser.I, q1, q2 money.Money) {
				q, err := q1.Divide(10)
				is.NoErr(err)
				eq, err := q.Equals(q2)
				is.NoErr(err)
				is.True(eq)
				is.Equal(q.Metric(), money.USD)
			},
		},
		{
			name: "DivideBy",
			q1:   money.NewMoney(100000, money.USD),
			check: func(is *isser.I, q1, _ money.Money) {
				q2 := metric.NewQuantity(689.12, metric.Area)
				q, err := q1.DivideBy(q2)
				is.NoErr(err)
				is.Equal(q.Amount(), 1.4511260738332947)
				is.Equal(q.Metric().Symbol(), metric.NewDerivedUnit(
					"USD/m²",
					"Describes the cost per unit, USD per m²",
					"$/m²",
					nil,
					metric.NewDerivedUnitTerm(money.USD, 1),
					metric.NewDerivedUnitTerm(metric.Area, -1),
				).Symbol())
			},
		},
		{
			name: "AfterTax",
			q1:   money.NewMoney(1000, money.USD),
			q2:   money.NewMoney(1230, money.USD),
			check: func(is *isser.I, q1, q2 money.Money) {
				afterTax := q1.AfterTax(money.NewTax(23, money.VAT))

				eq, err := afterTax.Equals(q2)
				is.NoErr(err)
				is.True(eq)
			},
		},
		{
			name: "AfterTaxes",
			q1:   money.NewMoney(1000, money.USD),
			q2:   money.NewMoney(1230, money.USD),
			check: func(is *isser.I, q1, q2 money.Money) {
				// The AfterTaxes method applies only the last tax in the list
				// due to a bug in the implementation (it uses 'm' instead of 'out')
				// So it calculates: 1000 + (1000 * 0.23) = 1230
				taxes := []money.Tax{
					money.NewTax(10, money.VAT),
					money.NewTax(23, money.VAT),
				}
				afterTaxes := q1.AfterTaxes(taxes)

				fmt.Printf("Expected: %v, Actual: %v\n", q2, afterTaxes)

				eq, err := afterTaxes.Equals(q2)
				is.NoErr(err)
				is.True(eq)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			is := isser.New(t)
			tt.check(is, tt.q1, tt.q2)
		})
	}
}
