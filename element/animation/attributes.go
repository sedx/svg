package animation

import (
	"encoding/xml"
	"fmt"

	"github.com/sedx/svg/types"
)

type restartAnimation string

const (
	Always        restartAnimation = "always"
	WhenNotActive                  = "whenNotActive"
	Never                          = "never"
)

type AnimationBegin interface {
	GetBeginValue() string
	MarshalXMLAttr(name xml.Name) (xml.Attr, error)
}

type TimigAttributes struct {
	Begin   AnimationBegin   `xml:"begin,attr,omitempty"`
	Dur     Duration         `xml:"dur,attr,omitempty"`
	Fill    types.FillType   `xml:"fill,attr,omitempty"`
	Restart restartAnimation `xml:"restart,attr,omitempty"`
	//  ‘end’, ‘min’, ‘max’, ‘restart’, ‘repeatCount’, ‘repeatDur’, ‘fill’
}

func (t *TimigAttributes) RestartAlways() {
	t.Restart = Always
}

func (t *TimigAttributes) RestartWhenNotActive() {
	t.Restart = WhenNotActive
}

func (t *TimigAttributes) RestartNever() {
	t.Restart = Never
}

type ValueAttributes struct {
	KeySplines Ease        `xml:"keySplines,attr,omitempty"`
	From       interface{} `xml:"from,attr,omitempty"`
	To         interface{} `xml:"to,attr,omitempty"`
	// ‘calcMode’, ‘values’, ‘keyTimes’, ‘keySplines’, ‘from’, ‘to’, ‘by’

}

type Ease struct {
	X1, Y1, X2, Y2 float64
}

func (e Ease) MarshalXMLAttr(n xml.Name) (xml.Attr, error) {
	if e.X1 == 0 && e.X2 == 0 && e.Y1 == 0 && e.Y2 == 0 {
		return xml.Attr{}, nil
	}
	return xml.Attr{Name: n, Value: fmt.Sprintf("%f, %f, %f, %f", e.X1, e.Y1, e.X2, e.Y2)}, nil
}

var (
	EaseIn    = Ease{0.42, 0, 1, 1}
	EaseOut   = Ease{0, 0, 0.58, 1}
	EaseInOut = Ease{0.42, 0, 0.58, 1}
)
