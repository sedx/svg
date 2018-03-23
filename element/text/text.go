package text

import (
	"encoding/xml"

	"github.com/sedx/svg/attributes"
)

// Text svg element
type Text struct {
	attributes.Core
	attributes.Positioned
	Content string   `xml:",chardata"`
	XMLName xml.Name `xml:"text"`
}
