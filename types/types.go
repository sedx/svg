package types

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/sedx/svg/unit"
)

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
	return xml.Attr{Name: name, Value: c.String()}, nil
}

// AnimationDuration
type AnimationDuration time.Duration

func (d AnimationDuration) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	val := float64(d) / float64(time.Second)
	return xml.Attr{Name: name, Value: fmt.Sprint(val)}, nil
}

type FillType string
