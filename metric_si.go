package metric

// SIBaseUnit represents a base unit of the International System of Units (SI).
type SIBaseUnit interface {
	Unit

	si() // marker method
}

var (
	SISystemOfUnits = NewSystemOfUnits("SI", "BIPM")

	Meter = newSIBaseUnit(
		"meter",
		"The meter is the length of the path travelled by light in vacuum during a time interval of 1/299792458 of a second",
		"m",
		SISystemOfUnits,
	)
	Radian = NewDerivedUnit(
		"radian",
		"The radian is the SI unit for measuring angles, and is the standard unit of angular measure used in many areas of mathematics",
		"rad",
		SISystemOfUnits,
	)
	Steradian = NewDerivedUnit(
		"steradian",
		"The steradian is the SI unit of solid angle. It is used to describe two-dimensional angles, analogous to the way in which the radian describes angles in three dimensions",
		"sr",
		SISystemOfUnits,
	)
	Area = NewDerivedUnit(
		"area",
		"The area is the quantity that expresses the extent of a two-dimensional figure or shape, or planar lamina, in the plane",
		"m²",
		SISystemOfUnits,
		NewDerivedUnitTerm(Meter, 2),
	)
	Volume = NewDerivedUnit(
		"volume",
		"The volume is the quantity of three-dimensional space enclosed by a closed surface, for example, the space that a substance (solid, liquid, gas, or plasma) or shape occupies or contains",
		"m³",
		SISystemOfUnits,
		NewDerivedUnitTerm(Meter, 3),
	)
	Kilogram = newSIBaseUnit(
		"kilogram",
		"The kilogram is the unit of mass; it is equal to the mass of the international prototype of the kilogram",
		"kg",
		SISystemOfUnits,
	)
	Second = newSIBaseUnit(
		"second",
		"The second is the duration of 9192631770 periods of the radiation corresponding to the transition between the two hyperfine levels of the ground state of the caesium 133 atom",
		"s",
		SISystemOfUnits,
	)
	Speed = NewDerivedUnit(
		"speed",
		"The speed is the rate of change of distance with time",
		"m/s",
		SISystemOfUnits,
		NewDerivedUnitTerm(Meter, 1),
		NewDerivedUnitTerm(Second, -1),
	)
	Ampere = newSIBaseUnit(
		"ampere",
		"The ampere is that constant current which, if maintained in two straight parallel conductors of infinite length, of negligible circular cross-section, and placed 1 meter apart in vacuum, would produce between these conductors a force equal to 2 x 10-7 newton per meter of length",
		"A",
		SISystemOfUnits,
	)
	Watt = NewDerivedUnit(
		"watt",
		"The watt is the SI derived unit for power in the International System of Units (SI); it is defined as 1 joule per second and is used to quantify the rate of energy transfer",
		"W",
		SISystemOfUnits,
		NewDerivedUnitTerm(Meter, 2),
		NewDerivedUnitTerm(Kilogram, 1),
		NewDerivedUnitTerm(Second, -3),
	)
	Kelvin = newSIBaseUnit(
		"kelvin",
		"The kelvin, unit of thermodynamic temperature, is the fraction 1/273.16 of the thermodynamic temperature of the triple point of water",
		"K",
		SISystemOfUnits,
	)
	Celsius = NewDerivedUnit(
		"celsius",
		"The degree Celsius is the unit of temperature defined by the equation T(°C) = T(K) - 273.15",
		"°C",
		SISystemOfUnits,
		NewDerivedUnitTerm(Kelvin, 1),
	)
	Mole = newSIBaseUnit(
		"mole",
		"The mole is the amount of substance of a system which contains as many elementary entities as there are atoms in 0.012 kilogram of carbon 12",
		"mol",
		SISystemOfUnits,
	)
	Candela = newSIBaseUnit(
		"candela",
		"The candela is the luminous intensity, in a given direction, of a source that emits monochromatic radiation of frequency 540 x 1012 hertz and that has a radiant intensity in that direction of 1/683 watt per steradian",
		"cd",
		SISystemOfUnits,
	)
	Lumen = NewDerivedUnit(
		"lumen",
		"The lumen is the SI derived unit of luminous flux, a measure of the total quantity of visible light emitted by a source per unit of time",
		"lm",
		SISystemOfUnits,
		NewDerivedUnitTerm(Candela, 1),
		NewDerivedUnitTerm(Steradian, 1),
	)
	Lux = NewDerivedUnit(
		"lux",
		"The lux is the SI unit of illuminance and luminous emittance, measuring luminous flux per unit area",
		"lx",
		SISystemOfUnits,
		NewDerivedUnitTerm(Candela, 1),
		NewDerivedUnitTerm(Meter, -2),
	)
)

type siBaseUnitImpl struct {
	name          string
	definition    string
	symbol        string
	systemOfUnits SystemOfUnits
}

func newSIBaseUnit(name, definition, symbol string, systemOfUnits SystemOfUnits) SIBaseUnit {
	unit := &siBaseUnitImpl{
		name:          name,
		definition:    definition,
		symbol:        symbol,
		systemOfUnits: systemOfUnits,
	}
	systemOfUnits.appendUnit(unit)

	return unit
}

func (s *siBaseUnitImpl) Name() string {
	return s.name
}

func (s *siBaseUnitImpl) Definition() string {
	return s.definition
}

func (s *siBaseUnitImpl) Symbol() string {
	return s.symbol
}

func (s *siBaseUnitImpl) SystemOfUnits() SystemOfUnits {
	return s.systemOfUnits
}

func (s *siBaseUnitImpl) String() string {
	return s.Symbol()
}

func (s *siBaseUnitImpl) si() {}
