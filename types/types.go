package types

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	"github.com/sedx/svg/unit"
)

const inherit = "inherit"

// XMLSpace is value of `xml:space` attributes
type XMLSpace string

// XMLSpace enum
const (
	XMLSpaceDefault  = "default"
	XMLSpacePreserve = "preserve"
)

// Length is SVG length type
type Length struct {
	Value interface{}
	Unit  unit.LengthUnit
}

func (l Length) String() string {
	return fmt.Sprintf("%v%s", l.Value, l.Unit)
}

// MarshalXMLAttr format length to attribute value
func (l Length) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if l.Value == nil {
		return xml.Attr{}, nil
	}
	return xml.Attr{Name: name, Value: l.String()}, nil
}

// Coordinate is SVG coordinate type
type Coordinate struct {
	Value interface{}
	Unit  unit.LengthUnit
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%v%s", c.Value, c.Unit)
}

// MarshalXMLAttr format coordinate to attribute value
func (c Coordinate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if c.Value == nil {
		return xml.Attr{}, nil
	}
	return xml.Attr{Name: name, Value: c.String()}, nil
}

// AnimationDuration
type AnimationDuration time.Duration

func (d AnimationDuration) GetBeginValue() string {
	return fmt.Sprint(float64(d) / float64(time.Second))
}

func (d AnimationDuration) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: d.GetBeginValue()}, nil
}

type FillType string

type Linecap string

const (
	Butt    Linecap = "butt"
	Round           = "round"
	Square          = "square"
	Inherit         = VectorEffect(inherit)
)

type VectorEffect string

const (
	NonScalingStroke VectorEffect = "non-scaling-stroke"
	None                          = "none"
)

type Transformations []Transform

func (t Transformations) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if len(t) == 0 {
		return xml.Attr{}, nil
	}
	var transformations []string
	for _, tr := range t {
		transformations = append(transformations, tr.GetTransform())
	}
	return xml.Attr{Name: name, Value: strings.Join(transformations, " ")}, nil
}

type Transform interface {
	GetTransform() string
}

type Color string

func (c Color) GetFill() string {
	return string(c)
}

func (c Color) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	var color string
	if color = c.GetFill(); color == "" {
		return xml.Attr{}, nil
	}
	return xml.Attr{Name: name, Value: color}, nil
}
