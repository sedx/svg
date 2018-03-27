package animation

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/sedx/svg/element"
	"github.com/sedx/svg/types"
)

type BeginEvent struct {
	Element element.Element
	Event   string
	Delay   time.Duration
}

func (b BeginEvent) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if b.Event == "" {
		return xml.Attr{}, nil
	}
	return xml.Attr{Name: name, Value: b.GetBeginValue()}, nil
}

func (b BeginEvent) GetBeginValue() string {
	return fmt.Sprintf("%s.%s+%s", b.Element.GetID(), b.Event, types.AnimationDuration(b.Delay).GetBeginValue())
}
