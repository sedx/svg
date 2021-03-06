package path

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/sedx/svg/element"

	"github.com/sedx/svg/attributes"
	"github.com/sedx/svg/element/animation"
	"github.com/sedx/svg/types"
)

// Path is SVG `path` element
type Path struct {
	element.SVGElement
	attributes.Stroked
	attributes.Filled
	D Data `xml:"d,attr,omitempty"`
	animation.Animated
	XMLName xml.Name `xml:"path"`
}

func (p *Path) addAction(a Action) {
	p.D = append(p.D, a)
}

// Data is slice of path actions (d attribute)
type Data []Action

func (d Data) String() string {
	var actions []string
	for _, action := range d {
		actions = append(actions, action.String())
	}
	return strings.Join(actions, " ")
}

// ActionParams is params for path command
type ActionParams []Param

type Param []interface{}

func (p Param) String() string {
	var params []string
	for _, param := range p {
		params = append(params, fmt.Sprint(param))
	}
	return strings.Join(params, ",")
}

// Action is one piece of path data
type Action struct {
	Command Command
	Params  ActionParams
}

func (a Action) String() string {
	var args []string
	for _, param := range a.Params {
		args = append(args, param.String())
	}
	return fmt.Sprintf("%s %s", a.Command, strings.Join(args, " "))
}

// M is SVG `moveto` path command
func (p *Path) M(x, y types.Coordinate) *Path {
	p.addAction(Action{M, ActionParams{Param{x, y}}})
	return p
}

func (p *Path) MR(x, y types.Coordinate) *Path {
	p.addAction(Action{MR, ActionParams{Param{x, y}}})
	return p
}

// L is SVG `lineto` path command
func (p *Path) L(x, y types.Coordinate) *Path {
	p.addAction(Action{L, ActionParams{Param{x, y}}})
	return p
}

// LR is SVG `lineto` path command (relative version)
func (p *Path) LR(x, y types.Coordinate) *Path {
	p.addAction(Action{LR, ActionParams{Param{x, y}}})
	return p
}

// Q is SVG `quadratic Bézier` curve commands
func (p *Path) Q(controlX, controlY, x, y types.Coordinate) *Path {
	p.addAction(Action{Q, ActionParams{Param{controlX, controlY}, Param{x, y}}})
	return p
}

// T is SVG command for shorthand/smooth quadratic Bézier curveto
func (p *Path) T(x, y types.Coordinate) *Path {
	p.addAction(Action{T, ActionParams{Param{x, y}}})
	return p
}

// C is SVG command for shorthand/smooth quadratic Bézier curveto
func (p *Path) C(controlX1, controlY1, controlX2, controlY2, x, y types.Coordinate) *Path {
	p.addAction(Action{C,
		ActionParams{
			Param{controlX1, controlY1},
			Param{controlX2, controlY2},
			Param{x, y},
		}})
	return p
}

//
func (p *Path) CR(controlX1, controlY1, controlX2, controlY2, x, y types.Coordinate) *Path {
	p.addAction(Action{CR,
		ActionParams{
			Param{controlX1, controlY1},
			Param{controlX2, controlY2},
			Param{x, y},
		}})
	return p
}

// S is SVG command for shorthand/smooth quadratic Bézier curveto
func (p *Path) S(controlX2, controlY2, x, y types.Coordinate) *Path {
	p.addAction(Action{S,
		ActionParams{
			Param{controlX2, controlY2},
			Param{x, y},
		}})
	return p
}

// S is SVG command for shorthand/smooth quadratic Bézier curveto (relative version)
func (p *Path) SR(controlX2, controlY2, x, y types.Coordinate) *Path {
	p.addAction(Action{SR,
		ActionParams{
			Param{controlX2, controlY2},
			Param{x, y},
		}})
	return p
}

// A  is SVG `A` path command
func (p *Path) A(rx, ry types.Coordinate, xAxisRotation float64, largeArcFlag, sweepFlag bool, x, y types.Coordinate) *Path {
	return p.arcCmd(false, rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y)
}

// AR  is SVG `a` path command
func (p *Path) AR(rx, ry types.Coordinate, xAxisRotation float64, largeArcFlag, sweepFlag bool, x, y types.Coordinate) *Path {
	return p.arcCmd(true, rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y)
}

func (p *Path) arcCmd(isRelative bool, rx, ry types.Coordinate, xAxisRotation float64, largeArcFlag, sweepFlag bool, x, y types.Coordinate) *Path {
	var isLarge int
	var isSweep int
	var cmd Command
	if largeArcFlag {
		isLarge = 1
	}
	if sweepFlag {
		isSweep = 1
	}
	cmd = A
	if isRelative {
		cmd = AR
	}
	p.addAction(Action{cmd,
		ActionParams{
			Param{rx, ry},
			Param{xAxisRotation},
			Param{isLarge, isSweep},
			Param{x, y},
		}})
	return p
}

// Z  is SVG `closepath` path command
func (p *Path) Z() *Path {
	p.addAction(Action{Command: Z})
	return p
}

// MarshalXMLAttr format length to attribute value
func (d Data) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if len(d) == 0 {
		return xml.Attr{}, nil
	}

	return xml.Attr{Name: name, Value: d.String()}, nil
}
