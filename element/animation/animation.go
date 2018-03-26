package animation

import (
	"encoding/xml"
	"time"

	"github.com/sedx/svg/attributes"
	"github.com/sedx/svg/types"
)

//
type Animate struct {
	attributes.AnimationTimed
	attributes.ValueAnimated
	AttributeType AttributeType `xml:"attributeType,attr,omitempty"`
	AttributeName string        `xml:"attributeName,attr,omitempty"`
	XMLName       xml.Name      `xml:"animate"`
}

func (a Animate) animate() {}

type AnimationElement interface {
	animate()
}

type Animated struct {
	Animations []AnimationElement `xml:"anim"`
}

type AttributeType string

const (
	CSS  AttributeType = "CSS"
	XML                = "XML"
	Auto               = "auto"
)

const (
	Freeze types.FillType = "freeze"
	Remove                = "remove"
)

func (a *Animated) AnimateCSS(attributeName string, from, to interface{}, dur time.Duration) *Animate {
	anim := &Animate{
		AttributeType: CSS,
		AttributeName: attributeName,
	}
	anim.From = from
	anim.To = to
	anim.Dur = types.AnimationDuration(dur)
	a.Animations = append(a.Animations, anim)
	return anim
}
