package shape

import (
	"encoding/xml"

	"github.com/sedx/svg/element"
	"github.com/sedx/svg/types"
)

// Line SVG element
type Line struct {
	element.SVGElement
	X1      types.Coordinate `xml:"x1,attr"`
	X2      types.Coordinate `xml:"x2,attr"`
	Y1      types.Coordinate `xml:"y1,attr"`
	Y2      types.Coordinate `xml:"y2,attr"`
	XMLName xml.Name         `xml:"line"`
}
