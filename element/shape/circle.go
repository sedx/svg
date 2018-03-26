package shape

import (
	"encoding/xml"

	"github.com/sedx/svg/element"

	"github.com/sedx/svg/attributes"
	"github.com/sedx/svg/types"
)

// Circle SVG element
type Circle struct {
	element.SVGElement
	attributes.Centred
	R       types.Length `xml:"r,attr"`
	XMLName xml.Name     `xml:"circle"`
}
