package animation

import (
	"encoding/xml"
)

type AnimateTransformType string

const (
	Translate AnimateTransformType = "translate"
	Scale                          = "scale"
	Rotate                         = "rotate"
	SkewX                          = "skewX"
	SkewY                          = "skewY"
)

type AnimateTransform struct {
	Animate
	Type    AnimateTransformType `xml:"type,attr"`
	XMLName xml.Name             `xml:"animateTransform"`
}

func (a AnimateTransform) animate() {}
