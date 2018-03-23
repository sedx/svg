package attributes

import (
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

type AnimationTimed struct {
	Begin types.AnimationDuration `xml:"begin,attr,omitempty"`
	Dur   types.AnimationDuration `xml:"dur,attr,omitempty"`
	Fill  types.FillType          `xml:"fill,attr,omitempty"`
	//  ‘end’, ‘min’, ‘max’, ‘restart’, ‘repeatCount’, ‘repeatDur’, ‘fill’
}

type ValueAnimated struct {
	From interface{} `xml:"from,attr,omitempty"`
	To   interface{} `xml:"to,attr,omitempty"`
	// ‘calcMode’, ‘values’, ‘keyTimes’, ‘keySplines’, ‘from’, ‘to’, ‘by’
}
