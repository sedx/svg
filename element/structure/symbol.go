package structure

import (
	"encoding/xml"

	"github.com/sedx/svg/element"
)

type Symbol struct {
	element.SVGElement
	element.SVGContainer
	XMLName xml.Name `xml:"symbol"`
}
