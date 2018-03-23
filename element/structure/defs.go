package structure

import (
	"encoding/xml"

	"github.com/sedx/svg/element"

	"github.com/sedx/svg/attributes"
)

type Defs struct {
	attributes.Core
	element.SVGContainer
	XMLName xml.Name `xml:"defs"`
}
