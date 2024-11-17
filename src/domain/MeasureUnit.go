package domain

type MeasureUnit int

const (
	Milligram  MeasureUnit = iota
	Gram       MeasureUnit = iota
	Kilogram   MeasureUnit = iota
	Teaspoon   MeasureUnit = iota
	Tablespoon MeasureUnit = iota
	Cup        MeasureUnit = iota
	Ounce      MeasureUnit = iota
	MilliLiter MeasureUnit = iota
	Liter      MeasureUnit = iota
)
