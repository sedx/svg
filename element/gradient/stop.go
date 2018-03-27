package gradient

import (
	"encoding/xml"

	"github.com/sedx/svg/attributes"
	"github.com/sedx/svg/element"
	"github.com/sedx/svg/element/animation"
	"github.com/sedx/svg/types"
)

type Stop struct {
	element.SVGElement
	animation.Animated
	attributes.Stoped
	Offset  types.Length `xml:"offset,attr,omitempty"`
	XMLName xml.Name     `xml:"stop"`
}
