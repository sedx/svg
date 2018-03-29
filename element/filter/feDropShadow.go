package filter

import (
	"encoding/xml"

	"github.com/sedx/svg/attributes"
	"github.com/sedx/svg/types"
)

type FeDropShadow struct {
	attributes.Core
	Primitive
	StdDeviation float64     `xml:"stdDeviation,attr,omitempty"`
	FloodColor   types.Color `xml:"flood-color,attr,omitempty"`
	FloodOpacity float64     `xml:"flood-opacity,attr,omitempty"`
	XMLName      xml.Name    `xml:"feDropShadow"`
}
