package attributes

import (
	"github.com/sedx/svg/element"
	"github.com/sedx/svg/types"
)

// Core attributes for all SVG element
type Core struct {
	ID    string         `xml:"id,attr,omitempty"`
	Base  string         `xml:"base,attr,omitempty"`
	Lang  string         `xml:"lang,attr,omitempty"`
	Space types.XMLSpace `xml:"space,attr,omitempty"`
	Class string         `xml:"class,attr,omitempty"`
	Style string         `xml:"style,attr,omitempty"`
	element.SVGElement
}

func (c Core) element() string {
	return c.ID
}
