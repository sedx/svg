package text

import (
	"encoding/xml"

	"github.com/sedx/svg/attributes"
	"github.com/sedx/svg/element"
	"github.com/sedx/svg/types"
)

// Text svg element
type Text struct {
	element.SVGElement
	attributes.Positioned
	attributes.Filled
	Dx      types.Length `xml:"dx,attr,omitempty"`
	Dy      types.Length `xml:"dy,attr,omitempty"`
	Content string       `xml:",chardata"`
	XMLName xml.Name     `xml:"text"`
}
