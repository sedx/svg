package shape

import (
	"encoding/xml"

	"github.com/sedx/svg/attributes"
	"github.com/sedx/svg/types"
)

// Circle SVG element
type Circle struct {
	attributes.Core
	attributes.Centred
	R       types.Length `xml:"r,attr"`
	XMLName xml.Name     `xml:"circle"`
}
