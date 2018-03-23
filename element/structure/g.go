package structure

import (
	"encoding/xml"

	"github.com/sedx/svg/attributes"
	"github.com/sedx/svg/element"
)

type G struct {
	attributes.Core
	element.SVGContainer
	XMLName xml.Name `xml:"g"`
}
