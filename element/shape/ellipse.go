package shape

import (
	"encoding/xml"

	"github.com/sedx/svg/attributes"
	"github.com/sedx/svg/element"
)

// Ellipse SVG element
type Ellipse struct {
	element.SVGElement
	attributes.Centred
	attributes.Rounded
	XMLName xml.Name `xml:"ellipse"`
}
