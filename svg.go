package svg

import (
	"encoding/xml"
	"io"
	"net/http"

	"github.com/sedx/svg/attributes"
	"github.com/sedx/svg/unit"

	"github.com/sedx/svg/element"
	"github.com/sedx/svg/element/path"
	"github.com/sedx/svg/element/shape"
	"github.com/sedx/svg/element/structure"
	"github.com/sedx/svg/element/style"
	"github.com/sedx/svg/element/text"
	"github.com/sedx/svg/types"
)

// SVG is svg document
type SVG struct {
	attributes.Measured
	element.SVGContainer
	// TODO make struct
	ViewBox             string   `xml:"viewBox,attr"`
	PreserveAspectRatio string   `xml:"preserveAspectRatio,attr"`
	XMLName             xml.Name `xml:"svg"`
}

type svgForMarshal struct {
	*SVG
	Version string `xml:"version,attr"`
	Ns      string `xml:"xmlns,attr"`
	Link    string `xml:"xlink,attr"`
}

// Marshal convert SVG document to byte slice
func (s *SVG) Marshal() ([]byte, error) {
	return xml.Marshal(s.prepare())
}

// MarshalIndent onvert SVG document to idented byte slice
func (s *SVG) MarshalIndent() ([]byte, error) {
	return xml.MarshalIndent(s.prepare(), "", " ")
}

func (s *SVG) prepare() *svgForMarshal {
	return &svgForMarshal{
		SVG:     s,
		Version: "1.1",
		Ns:      "http://www.w3.org/2000/svg",
		Link:    "http://www.w3.org/1999/xlink",
	}
}

// Length return svg length type
func Length(v interface{}, u ...unit.LengthUnit) types.Length {
	var unit unit.LengthUnit
	if len(u) > 0 {
		unit = u[0]
	}
	return types.Length{Value: v, Unit: unit}
}

// Coordinate return svg length type
func Coordinate(v interface{}, u ...unit.LengthUnit) types.Coordinate {
	var unit unit.LengthUnit
	if len(u) > 0 {
		unit = u[0]
	}
	return types.Coordinate{Value: v, Unit: unit}
}

// Text return `text` SVG element
func Text(x, y types.Coordinate, content string) *text.Text {
	t := &text.Text{Content: content}
	t.X = x
	t.Y = y
	return t
}

// G return `g` SVG element
func G() *structure.G {
	return &structure.G{}
}

func Defs() *structure.Defs {
	return &structure.Defs{}
}

// Circle return `circle` SVG element
func Circle(x, y types.Coordinate, r types.Length) *shape.Circle {
	c := &shape.Circle{R: r}
	c.X = x
	c.Y = y
	return c
}

// Ellipse return `ellipse` SVG element
func Ellipse(x, y types.Coordinate, rx types.Length, ry ...types.Length) *shape.Ellipse {
	e := &shape.Ellipse{}
	e.X = x
	e.Y = y
	e.Rx = rx
	if len(ry) > 0 {
		e.Ry = ry[0]
	} else {
		e.Ry = rx
	}
	return e
}

// Rect return `rect` SVG element
func Rect(x, y types.Coordinate) *shape.Rect {
	r := &shape.Rect{}
	r.X = x
	r.Y = y
	return r
}

// RoundRect return `rect` SVG element with rounded corners
func RoundRect(x, y types.Coordinate, roundX types.Length, roundY ...types.Length) *shape.Rect {
	r := Rect(x, y)
	r.Rx = roundX
	if len(roundY) > 0 {
		r.Ry = roundY[0]
	} else {
		r.Ry = roundX
	}
	return r
}

// Line return `line` SVG element
func Line(x1, x2, y1, y2 types.Coordinate) *shape.Line {
	return &shape.Line{X1: x1, X2: x2, Y1: y1, Y2: y2}
}

func Path() *path.Path {
	return &path.Path{}
}

// Struct interface realized by svg struct elements
// that can add nested elements
type Struct interface {
	Add(e element.Element)
}

func (s *SVG) Serve(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "image/svg+xml")
	b, err := s.MarshalIndent()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
	}
	w.Write(b)
}

func Style(styles string) *style.Style {
	return &style.Style{Content: styles}
}
