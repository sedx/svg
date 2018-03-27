package animation

import (
	"encoding/xml"
)

type Set struct {
	TimigAttributes
	To            interface{}   `xml:"to,attr,omitempty"`
	AttributeType AttributeType `xml:"attributeType,attr,omitempty"`
	AttributeName string        `xml:"attributeName,attr,omitempty"`
	XMLName       xml.Name      `xml:"set"`
}

func (s Set) animate() {}
