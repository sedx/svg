package attributes

import (
	"encoding/xml"

	"github.com/sedx/svg/types"
)

// Positioned attributes for element that have x and y
type Positioned struct {
	X types.Coordinate `xml:"x,attr"`
	Y types.Coordinate `xml:"y,attr"`
}

// Measured attributes for element that have width and heighth
type Measured struct {
	Width  types.Length `xml:"width,attr,omitempty"`
	Height types.Length `xml:"height,attr,omitempty"`
}

// Rounded attributes for element that has rx and ry
type Rounded struct {
	Rx types.Length `xml:"rx,attr,omitempty"`
	Ry types.Length `xml:"ry,attr,omitempty"`
}

// Centred attributes that have cx anf cy
type Centred struct {
	X types.Coordinate `xml:"cx,attr"`
	Y types.Coordinate `xml:"cy,attr"`
}

type ValueAnimated struct {
	From interface{} `xml:"from,attr,omitempty"`
	To   interface{} `xml:"to,attr,omitempty"`
	// ‘calcMode’, ‘values’, ‘keyTimes’, ‘keySplines’, ‘from’, ‘to’, ‘by’
}

type Stroked struct {
	Stroke        string             `xml:"stroke,attr,omitempty"`
	StrokeLinecap types.Linecap      `xml:"stroke-linecap,attr,omitempty"`
	StrokeWidth   types.Length       `xml:"stroke-width,attr,omitempty"`
	VectorEffect  types.VectorEffect `xml:"vector-effect,attr,omitempty"`
}

type Filling interface {
	GetFill() string
	MarshalXMLAttr(xml.Name) (xml.Attr, error)
}

type Filled struct {
	Fill Filling `xml:"fill,attr,omitempty"`
}
}

type ViewBoxed struct {
	// TODO make struct
	ViewBox             string `xml:"viewBox,attr,omitempty"`
	PreserveAspectRatio string `xml:"preserveAspectRatio,attr,omitempty"`
}

type Transformable struct {
	Transforms types.Transformations `xml:"transform,attr,omitempty"`
}

func (t *Transformable) Transform(transformation types.Transform) {
	t.Transforms = append(t.Transforms, transformation)
}
