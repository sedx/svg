package structure

import (
	"encoding/xml"

	"github.com/sedx/svg/element"
)

type Defs struct {
	element.SVGElement
	element.SVGContainer
	XMLName xml.Name `xml:"defs"`
}
