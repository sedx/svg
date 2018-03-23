package shape

import (
	"encoding/xml"

	"github.com/sedx/svg/attributes"
)

// Ellipse SVG element
type Ellipse struct {
	attributes.Core
	attributes.Centred
	attributes.Rounded
	XMLName xml.Name `xml:"ellipse"`
}
