package metric

// Metric describes a standard of measurement.
type Metric interface {
	// Name returns the name of the metric. For example, "weight".
	Name() string

	// Definition returns the formal definition of the Metric,
	// e.g., "The meter is the length of the path traveled by light in vacuum during a time interval of 1/299792458 of a second"
	Definition() string

	// Symbol returns "" or the standard symbol for the Metric, e.g., "m"
	Symbol() string

	String() string
}

type metricImpl struct {
	name       string
	definition string
	symbol     string
}

// NewMetric creates a new Metric instance.
func NewMetric(name, definition, symbol string) Metric {
	return &metricImpl{
		name:       name,
		definition: definition,
		symbol:     symbol,
	}
}

func (m *metricImpl) Name() string {
	return m.name
}

func (m *metricImpl) Definition() string {
	return m.definition
}

func (m *metricImpl) Symbol() string {
	return m.symbol
}

func (m *metricImpl) String() string {
	return m.Symbol()
}

// Unit describes a Metric that is a part of a SystemOfUnits.
type Unit interface {
	Metric

	SystemOfUnits() SystemOfUnits
}

// SystemOfUnits describes a set of related Units defined by a standardization body.
type SystemOfUnits interface {
	Name() string
	StandardizationBody() string

	Units() []Unit

	appendUnit(unit Unit)
}

type systemOfUnitsImpl struct {
	name                string
	standardizationBody string
	units               []Unit
}

func NewSystemOfUnits(name, standardizationBody string) SystemOfUnits {
	return &systemOfUnitsImpl{
		name:                name,
		standardizationBody: standardizationBody,
		units:               make([]Unit, 0),
	}
}

func (s *systemOfUnitsImpl) Name() string {
	return s.name
}

func (s *systemOfUnitsImpl) StandardizationBody() string {
	return s.standardizationBody
}

func (s *systemOfUnitsImpl) Units() []Unit {
	return s.units
}

func (s *systemOfUnitsImpl) appendUnit(unit Unit) {
	s.units = append(s.units, unit)
}
