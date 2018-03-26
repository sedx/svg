package style

import (
	"encoding/xml"

	"github.com/sedx/svg/element"
)

type Style struct {
	element.SVGElement
	Content string   `xml:",chardata"`
	XMLName xml.Name `xml:"style"`
}
