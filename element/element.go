package element

import (
	"encoding/xml"
	"fmt"

	"github.com/sedx/svg/attributes"
)

// Element is fake interface to filter elements
type Element interface {
	MarshalXMLAttr(name xml.Name) (xml.Attr, error)
	GetID() string
}

// SVGElement used for composition to realise Element interface
type SVGElement struct {
	attributes.Core
	Class string `xml:"class,attr,omitempty"`
	Style string `xml:"style,attr,omitempty"`
}

func (s SVGElement) GetID() string {
	return s.ID
}

func (s SVGElement) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if s.ID == "" {
		return xml.Attr{}, fmt.Errorf("Element can be linked to %s, cause ID is missed", name.Local)
	}
	return xml.Attr{Name: name, Value: fmt.Sprintf("#%s", s.ID)}, nil
}

// SVGContainer is holder of other SVG elements
type SVGContainer struct {
	Elements []Element
}

// Add append element to elements slice
func (s *SVGContainer) Add(el ...Element) {
	s.Elements = append(s.Elements, el...)
}
