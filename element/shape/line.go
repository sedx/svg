package shape

import (
	"encoding/xml"

	"github.com/sedx/svg/attributes"
	"github.com/sedx/svg/types"
)

// Line SVG element
type Line struct {
	attributes.Core
	X1      types.Coordinate `xml:"x1,attr"`
	X2      types.Coordinate `xml:"x2,attr"`
	Y1      types.Coordinate `xml:"y1,attr"`
	Y2      types.Coordinate `xml:"y2,attr"`
	XMLName xml.Name         `xml:"line"`
}
