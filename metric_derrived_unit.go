package metric

import (
	"fmt"
)

// DerivedUnit represents a derived unit composed of terms.
// It extends the Unit interface and provides a method to retrieve the terms.
type DerivedUnit interface {
	Unit

	Terms() []DerivedUnitTerm
}

type derivedUnitImpl struct {
	name          string
	definition    string
	symbol        string
	systemOfUnits SystemOfUnits
	terms         []DerivedUnitTerm
}

// NewDerivedUnit creates a new DerivedUnit with the given name, definition, symbol, systemOfUnits, and terms.
// The terms are the units that make up the DerivedUnit, e.g., the DerivedUnit "meter per second" has two terms: "meter" and "second".
// The systemOfUnits is the SystemOfUnits that the DerivedUnit belongs to (e.g., the SI SystemOfUnits).
// If the DerivedUnit does not belong to a SystemOfUnits, then pass nil.
func NewDerivedUnit(name, definition, symbol string, systemOfUnits SystemOfUnits, terms ...DerivedUnitTerm) DerivedUnit {
	du := &derivedUnitImpl{
		name:          name,
		definition:    definition,
		symbol:        symbol,
		systemOfUnits: systemOfUnits,
		terms:         terms,
	}

	if systemOfUnits != nil {
		systemOfUnits.appendUnit(du)
	}

	return du
}

func (d *derivedUnitImpl) Name() string {
	return d.name
}

func (d *derivedUnitImpl) String() string {
	return d.Symbol()
}

func (d *derivedUnitImpl) Definition() string {
	return d.definition
}

func (d *derivedUnitImpl) Symbol() string {
	return d.symbol
}

func (d *derivedUnitImpl) SystemOfUnits() SystemOfUnits {
	return d.systemOfUnits
}

func (d *derivedUnitImpl) Terms() []DerivedUnitTerm {
	return d.terms
}

type DerivedUnitTerm interface {
	Metric

	Exponent() int
	Metric() Metric
}

type derivedUnitTermImpl struct {
	metric   Metric
	exponent int
}

// NewDerivedUnitTerm creates a new DerivedUnitTerm with the given unit and exponent.
func NewDerivedUnitTerm(metric Metric, exponent int) DerivedUnitTerm {
	du := &derivedUnitTermImpl{
		metric:   metric,
		exponent: exponent,
	}
	return du
}

func (d *derivedUnitTermImpl) Name() string {
	return d.metric.Name()
}

func (d *derivedUnitTermImpl) String() string {
	return d.Symbol()
}

func (d *derivedUnitTermImpl) Definition() string {
	return d.metric.Definition()
}

func (d *derivedUnitTermImpl) Symbol() string {
	return d.metric.Symbol() + "^" + fmt.Sprint(d.exponent)
}

func (d *derivedUnitTermImpl) Exponent() int {
	return d.exponent
}

func (d *derivedUnitTermImpl) Metric() Metric {
	return d.metric
}
