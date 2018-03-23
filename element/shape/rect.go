package shape

import (
	"encoding/xml"

	"github.com/sedx/svg/attributes"
)

type Rect struct {
	attributes.Core
	attributes.Positioned
	attributes.Measured
	attributes.Rounded
	XMLName xml.Name `xml:"rect"`
}
