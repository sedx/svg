package shape

import (
	"encoding/xml"

	"github.com/sedx/svg/element/animation"

	"github.com/sedx/svg/attributes"
	"github.com/sedx/svg/element"
)

type Rect struct {
	element.SVGElement
	attributes.Positioned
	attributes.Measured
	attributes.Rounded
	attributes.Filled
	attributes.Stroked
	attributes.Transformable
	animation.Animated
	XMLName xml.Name `xml:"rect"`
}
