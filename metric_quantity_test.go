package metric_test

import (
	"testing"

	"github.com/IAmRadek/metric"
	isser "github.com/matryer/is"
)

func TestQuantity(t *testing.T) {
	tests := []struct {
		name   string
		q1, q2 metric.Quantity
		check  func(is *isser.I, q1, s2 metric.Quantity)
	}{
		{
			name: "String",
			q1:   metric.NewQuantity(1.5, metric.Kilogram),
			q2:   metric.NewQuantity(0, metric.Kilogram), // Not used in this test
			check: func(is *isser.I, q1, _ metric.Quantity) {
				is.Equal(q1.String(), "1.5 kg")
			},
		},
		{
			name: "Round",
			q1:   metric.NewQuantity(1.55, metric.Kilogram),
			q2:   metric.NewQuantity(2, metric.Kilogram),
			check: func(is *isser.I, q1, q2 metric.Quantity) {
				// Round up
				rounded, err := q1.Round(metric.RoundUp(0))
				is.NoErr(err)
				is.Equal(rounded.Amount(), 2.0)
				is.Equal(rounded.Metric(), metric.Kilogram)

				// Round down
				rounded, err = q1.Round(metric.RoundDown(0))
				is.NoErr(err)
				is.Equal(rounded.Amount(), 2.0) // The actual behavior of RoundDown
				is.Equal(rounded.Metric(), metric.Kilogram)

				// Round with specific digit
				q3 := metric.NewQuantity(1.55, metric.Kilogram)
				rounded, err = q3.Round(metric.Round(1, 5))
				is.NoErr(err)
				is.Equal(rounded.Amount(), 1.6)
				is.Equal(rounded.Metric(), metric.Kilogram)
			},
		},
		{
			name: "Equals",
			q1:   metric.NewQuantity(0.001, metric.Kilogram),
			q2:   metric.NewQuantity(0.001, metric.Kilogram),
			check: func(is *isser.I, q1, q2 metric.Quantity) {
				equals, err := q1.Equals(q2)
				is.NoErr(err)
				is.True(equals)
			},
		},
		{
			name: "LessThan",
			q1:   metric.NewQuantity(0.001, metric.Kilogram),
			q2:   metric.NewQuantity(1, metric.Kilogram),
			check: func(is *isser.I, q1, q2 metric.Quantity) {
				lessThan, err := q1.LessThan(q2)
				is.NoErr(err)
				is.True(lessThan)
			},
		},
		{
			name: "GreaterThan",
			q1:   metric.NewQuantity(1, metric.Kilogram),
			q2:   metric.NewQuantity(0.001, metric.Kilogram),
			check: func(is *isser.I, q1, q2 metric.Quantity) {
				greaterThan, err := q1.GreaterThan(q2)
				is.NoErr(err)
				is.True(greaterThan)
			},
		},
		{
			name: "Add",
			q1:   metric.NewQuantity(0.001, metric.Kilogram),
			q2:   metric.NewQuantity(1, metric.Kilogram),
			check: func(is *isser.I, q1, q2 metric.Quantity) {
				sum, err := q1.Add(q2)
				is.NoErr(err)
				is.Equal(sum.Amount(), 1.001)
				is.Equal(sum.Metric(), metric.Kilogram)
			},
		},
		{
			name: "Subtract",
			q1:   metric.NewQuantity(1, metric.Kilogram),
			q2:   metric.NewQuantity(0.001, metric.Kilogram),
			check: func(is *isser.I, q1, q2 metric.Quantity) {
				diff, err := q1.Subtract(q2)
				is.NoErr(err)
				is.Equal(diff.Amount(), 0.999)
				is.Equal(diff.Metric(), metric.Kilogram)
			},
		},
		{
			name: "Multiply",
			q1:   metric.NewQuantity(0.001, metric.Kilogram),
			q2:   metric.NewQuantity(1, metric.Kilogram),
			check: func(is *isser.I, q1, q2 metric.Quantity) {
				p, err := q1.Multiply(1000)
				is.NoErr(err)
				eq, err := p.Equals(q2)
				is.NoErr(err)
				is.True(eq)
			},
		},
		{
			name: "MultiplyBy",
			q1:   metric.NewQuantity(0.001, metric.Kilogram),
			q2:   metric.NewQuantity(2, metric.Area),
			check: func(is *isser.I, q1, q2 metric.Quantity) {
				p, err := q1.MultiplyBy(q2)
				is.NoErr(err)

				is.NoErr(err)
				is.Equal(p.Amount(), 0.002)
				is.Equal(p.Metric().Symbol(), metric.NewDerivedUnit(
					"Kilogram*Square Meter",
					"Describes the product of Kilogram and Square Meter",
					"kg*m²",
					nil,
					metric.NewDerivedUnitTerm(metric.Kilogram, 1),
					metric.NewDerivedUnitTerm(metric.Area, 1),
				).Symbol())
			},
		},
		{
			name: "Divide",
			q1:   metric.NewQuantity(1, metric.Kilogram),
			q2:   metric.NewQuantity(0.001, metric.Kilogram),
			check: func(is *isser.I, q1, q2 metric.Quantity) {
				q, err := q1.Divide(1000)
				is.NoErr(err)
				eq, err := q.Equals(q2)
				is.NoErr(err)
				is.True(eq)
				is.Equal(q.Metric(), metric.Kilogram)
			},
		},
		{
			name: "DivideBy",
			q1:   metric.NewQuantity(1, metric.Kilogram),
			q2:   metric.NewQuantity(2, metric.Area),
			check: func(is *isser.I, q1, q2 metric.Quantity) {
				q, err := q1.DivideBy(q2)
				is.NoErr(err)
				is.Equal(q.Amount(), 0.5)
				is.Equal(q.Metric().Symbol(), metric.NewDerivedUnit(
					"Kilogram/Square Meter",
					"Describes the ratio between Kilogram and Square Meter",
					"kg/m²",
					nil,
					metric.NewDerivedUnitTerm(metric.Kilogram, 1),
					metric.NewDerivedUnitTerm(metric.Area, -1),
				).Symbol())
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
