package text

import (
	"encoding/xml"

	"github.com/sedx/svg/attributes"
	"github.com/sedx/svg/element"
)

// Text svg element
type Text struct {
	element.SVGElement
	attributes.Positioned
	Content string   `xml:",chardata"`
	XMLName xml.Name `xml:"text"`
}
