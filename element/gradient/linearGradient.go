package gradient

import (
	"encoding/xml"
	"fmt"

	"github.com/sedx/svg/attributes"
	"github.com/sedx/svg/element"
	"github.com/sedx/svg/element/animation"
	"github.com/sedx/svg/types"
)

type gradient struct {
	element.SVGElement
	attributes.Stroked
	attributes.Filled
	animation.Animated
	Stops []*Stop
	// todo enum same as for mask
	GradientUnits string `xml:"gradientUnits,attr,omitempty"`
}

func (g *gradient) AddStop(s ...*Stop) {
	g.Stops = append(g.Stops, s...)
}

func (g gradient) GetFill() string {
	return fmt.Sprintf("url(#%s)", g.ID)
}

func (g gradient) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	var gradient string
	if gradient = g.GetFill(); gradient == "" {
		return xml.Attr{}, nil
	}
	return xml.Attr{Name: name, Value: gradient}, nil
}

type LinearGradient struct {
	gradient
	X1      types.Coordinate `xml:"x1,attr,omitempty"`
	Y1      types.Coordinate `xml:"y1,attr,omitempty"`
	X2      types.Coordinate `xml:"x2,attr,omitempty"`
	Y2      types.Coordinate `xml:"y2,attr,omitempty"`
	XMLName xml.Name         `xml:"linearGradient"`
}
