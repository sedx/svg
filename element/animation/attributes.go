package animation

import (
	"encoding/xml"

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
	Begin   AnimationBegin          `xml:"begin,attr,omitempty"`
	Dur     types.AnimationDuration `xml:"dur,attr,omitempty"`
	Fill    types.FillType          `xml:"fill,attr,omitempty"`
	Restart restartAnimation        `xml:"restart,attr,omitempty"`
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
