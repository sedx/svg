package unit

// Unit is base for allowed in svg units
type Unit string

// AngleUnit is SVG angle units
type AngleUnit Unit

// supported Angle units
const (
	Deg  AngleUnit = "deg"
	Grad           = "grad"
	Rad            = "rad"
)

// FrequencyUnit is SVG frequency units
type FrequencyUnit Unit

// supported Frequency units
const (
	Hz  FrequencyUnit = "Hz"
	KHz               = "kHz"
)

// LengthUnit is  SVG length units
type LengthUnit Unit

// supported length units
const (
	EM      LengthUnit = "em"
	EX                 = "ex"
	PX                 = "px"
	IN                 = "in"
	CM                 = "cm"
	MM                 = "mm"
	PT                 = "pt"
	PC                 = "pc"
	Percent            = "%"
)
