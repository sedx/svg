package element

import (
	"encoding/xml"
	"fmt"

	"github.com/sedx/svg/types"
)

// Element is fake interface to filter elements
type Element interface {
	MarshalXMLAttr(name xml.Name) (xml.Attr, error)
}

// SVGElement used for composition to realise Element interface
type SVGElement struct {
	ID    string         `xml:"id,attr,omitempty"`
	Base  string         `xml:"base,attr,omitempty"`
	Lang  string         `xml:"lang,attr,omitempty"`
	Space types.XMLSpace `xml:"space,attr,omitempty"`
	Class string         `xml:"class,attr,omitempty"`
	Style string         `xml:"style,attr,omitempty"`
}

// func (s SVGElement) element() {}

func (s SVGElement) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if s.ID == "" {
		return xml.Attr{}, fmt.Errorf("%+v element has no ID that required for linking", s)
	}
	return xml.Attr{Name: name, Value: fmt.Sprintf("#%s", s.ID)}, nil
}

// SVGContainer is holder of other SVG elements
type SVGContainer struct {
	Elements []Element
}

// Add append element to elements slice
func (s *SVGContainer) Add(e Element) {
	s.Elements = append(s.Elements, e)
}
