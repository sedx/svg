package structure

import (
	"encoding/xml"

	"github.com/sedx/svg/element/mask"

	"github.com/sedx/svg/attributes"
	"github.com/sedx/svg/element"
)

type Use struct {
	element.SVGContainer
	element.SVGElement
	attributes.Positioned
	attributes.Measured
	attributes.ViewBoxed
	Mask      *mask.Mask      `xml:"mask,attr,omitempty"`
	XLinkHref element.Element `xml:"xlink:href,attr"`
	XMLName   xml.Name        `xml:"use"`
}
