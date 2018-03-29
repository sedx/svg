package animation

import (
	"encoding/xml"
	"fmt"
	"time"
)

// AnimationDuration
type Duration time.Duration

func (d Duration) GetBeginValue() string {
	return fmt.Sprint(float64(d) / float64(time.Second))
}

func (d Duration) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: d.GetBeginValue()}, nil
}
