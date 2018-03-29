package filter

import (
	"encoding/xml"

	"github.com/sedx/svg/attributes"
)

type BlendMode string

const (
	BlendModeNormal   BlendMode = "normal"
	BlendModeMultiply           = "multiply"
	BlendModeScreen             = "screen"
	BlendModeDarken             = "darken"
	BlendModeLighten            = "lighten"
)

type FeBlend struct {
	attributes.Core
	Primitive
	In      FilterOut `xml:"in,attr"`
	In2     FilterOut `xml:"in2,attr"`
	Mode    BlendMode `xml:"mode,attr"`
	XMLName xml.Name  `xml:"feBlend"`
}
