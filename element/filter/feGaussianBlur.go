package filter

import (
	"encoding/xml"

	"github.com/sedx/svg/attributes"
)

type FeGaussianBlur struct {
	attributes.Core
	Primitive
	StdDeviation float64   `xml:"stdDeviation,attr,omitempty"`
	In           FilterOut `xml:"in,attr"`
	XMLName      xml.Name  `xml:"feGaussianBlur"`
}

// var f = FeGaussianBlur{}
// var x = f.M
