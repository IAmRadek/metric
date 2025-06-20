package money_test

import (
	"testing"

	"github.com/IAmRadek/metric/money"
	isser "github.com/matryer/is"
)

func TestCurrency(t *testing.T) {
	tests := []struct {
		name     string
		currency money.Currency
		check    func(is *isser.I, currency money.Currency)
	}{
		{
			name:     "Name",
			currency: money.USD,
			check: func(is *isser.I, currency money.Currency) {
				is.Equal(currency.Name(), "US Dollar")
			},
		},
		{
			name:     "Definition",
			currency: money.USD,
			check: func(is *isser.I, currency money.Currency) {
				is.Equal(currency.Definition(), "The US Dollar is the official currency of the United States and its territories per the Coinage Act of 1792. One dollar is divided into 100 cents (Symbol: ¢).")
			},
		},
		{
			name:     "NewNonISOCurrency",
			currency: money.NewNonISOCurrency("Bitcoin", "A decentralized digital currency", "₿", "BTC", 8),
			check: func(is *isser.I, currency money.Currency) {
				is.Equal(currency.Name(), "Bitcoin")
				is.Equal(currency.Definition(), "A decentralized digital currency")
				is.Equal(currency.Symbol(), "₿")
				is.Equal(currency.Code(), "BTC")
				is.Equal(currency.Decimal(), 8)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			is := isser.New(t)
			tt.check(is, tt.currency)
		})
	}
}

func TestISOCurrencies(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		is := isser.New(t)

		// Test getting an existing currency
		currency, ok := money.ISOCurrencies.Get("USD")
		is.True(ok)
		is.Equal(currency.Code(), "USD")

		// Test getting a non-existing currency
		_, ok = money.ISOCurrencies.Get("XYZ")
		is.True(!ok)
	})
}
