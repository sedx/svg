package script

import (
	"encoding/xml"

	"github.com/sedx/svg/attributes"

	"github.com/sedx/svg/element"
)

type Script struct {
	element.SVGElement
	attributes.Core
	Type    string   `xml:"type,attr,omitempty"`
	Content []byte   `xml:",cdata"`
	XMLName xml.Name `xml:"script"`
}
