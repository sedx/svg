package structure

import (
	"encoding/xml"

	"github.com/sedx/svg/element"
)

type G struct {
	element.SVGElement
	element.SVGContainer
	XMLName xml.Name `xml:"g"`
}
