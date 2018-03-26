package structure

import (
	"encoding/xml"

	"github.com/sedx/svg/element/animation"

	"github.com/sedx/svg/element"
)

type G struct {
	element.SVGElement
	element.SVGContainer
	animation.Animated
	XMLName xml.Name `xml:"g"`
}
