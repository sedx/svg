package mask

import (
	"encoding/xml"
	"fmt"

	"github.com/sedx/svg/attributes"
	"github.com/sedx/svg/element"
)

type Mask struct {
	element.SVGContainer
	element.SVGElement
	attributes.Positioned
	attributes.Measured
	// todo enum
	MaskContentUnits string   `xml:"maskContentUnits,attr,omitempty"`
	MaskUnits        string   `xml:"maskUnits,attr,omitempty"`
	XMLName          xml.Name `xml:"mask"`
}

func (m *Mask) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if m == nil {
		return xml.Attr{}, nil
	}
	if m.ID == "" {
		return xml.Attr{}, fmt.Errorf("ID is missing for %s", name)
	}
	return xml.Attr{Name: name, Value: fmt.Sprintf("url(#%s)", m.GetID())}, nil
}
