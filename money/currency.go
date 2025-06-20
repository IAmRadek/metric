package money

import (
	"github.com/IAmRadek/metric"
)

type Currency interface {
	metric.Metric

	// Code is an alphabetic code that represents the currency, e.g., "EUR" for the Euro
	Code() string

	// Decimal returns the number of decimal places for the currency, e.g., 2 for USD representing dollars and cents.
	Decimal() int
}

// ISOCurrency is a Currency defined by the ISO 4217 standard.
type ISOCurrency interface {
	Currency
}

type NonISOCurrency interface {
	Currency
}

type isoCurrencies interface {
	Get(code string) (ISOCurrency, bool)
}

type isoCurrenciesImpl struct {
	m map[string]ISOCurrency
}

func (i isoCurrenciesImpl) Get(code string) (ISOCurrency, bool) {
	c, ok := i.m[code]
	return c, ok
}

var ISOCurrencies isoCurrencies = isoCurrenciesImpl{
	m: make(map[string]ISOCurrency),
}
var (
	USD = NewISOCurrency(
		"US Dollar",
		"The US Dollar is the official currency of the United States and its territories per the Coinage Act of 1792. One dollar is divided into 100 cents (Symbol: ¢).",
		"$",
		"USD",
		2,
	)

	EUR = NewISOCurrency(
		"Euro",
		"The Euro is the official currency of the European Union. One euro is divided into 100 cents (Symbol: ¢).",
		"€",
		"EUR",
		2,
	)

	PLN = NewISOCurrency(
		"Polish Zloty",
		"The Polish Zloty is the official currency of Poland. One zloty is divided into 100 groszy (Symbol: gr).",
		"zł",
		"PLN",
		2,
	)

	GBP = NewISOCurrency(
		"Pound Sterling",
		"The Pound Sterling is the official currency of the United Kingdom and its territories. One pound sterling is divided into 100 pence (Symbol: p).",
		"£",
		"GBP",
		2,
	)
)

type currencyImpl struct {
	name       string
	definition string
	symbol     string
	code       string
	decimal    int
}

func NewCurrency(name, definition, symbol, code string, decimal int) Currency {
	return &currencyImpl{
		name:       name,
		definition: definition,
		symbol:     symbol,
		code:       code,
		decimal:    decimal,
	}
}

// NewISOCurrency creates a new ISOCurrency and adds it to the ISOCurrencies map.
func NewISOCurrency(name, definition, symbol, code string, decimal int) ISOCurrency {
	currency := NewCurrency(name, definition, symbol, code, decimal)

	ISOCurrencies.(isoCurrenciesImpl).m[code] = currency

	return currency
}

// NewNonISOCurrency creates a new NonISOCurrency instance by calling NewCurrency with the provided parameters.
func NewNonISOCurrency(name, definition, symbol, code string, decimal int) NonISOCurrency {
	return NewCurrency(name, definition, symbol, code, decimal)
}

func (c currencyImpl) Name() string {
	return c.name
}

func (c currencyImpl) Definition() string {
	return c.definition
}

func (c currencyImpl) Symbol() string {
	return c.symbol
}

func (c currencyImpl) Code() string {
	return c.code
}

func (c currencyImpl) Decimal() int {
	return c.decimal
}

func (c currencyImpl) String() string {
	return c.Symbol()
}
