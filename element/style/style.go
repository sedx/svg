package style

import (
	"encoding/xml"

	"github.com/sedx/svg/attributes"
)

type Style struct {
	attributes.Core
	Content string   `xml:",chardata"`
	XMLName xml.Name `xml:"style"`
}
