package filter

import (
	"encoding/xml"
	"fmt"

	"github.com/sedx/svg/attributes"
	"github.com/sedx/svg/element"
	"github.com/sedx/svg/types"
)

type Primitive struct {
	attributes.Filtered
}

func (f Primitive) MarshalXMLAttr(n xml.Name) (xml.Attr, error) {
	outName, err := f.GetOut()
	if err != nil {
		return xml.Attr{}, err
	}
	return xml.Attr{Name: n, Value: outName}, nil
}

type FilterPrimitive interface {
	FilterOut
}

type FilterOut interface {
	GetOut() (string, error)
	MarshalXMLAttr(xml.Name) (xml.Attr, error)
}

type Filter struct {
	element.SVGElement
	attributes.Measured
	attributes.Positioned
	FilterUnits    types.CoordinateSystemUnits `xml:"filterUnits,attr,omitempty"`
	PrimitiveUnits types.CoordinateSystemUnits `xml:"primitiveUnits,attr,omitempty"`
	Primitives     []FilterPrimitive
	XMLName        xml.Name `xml:"filter"`
}

func (f Filter) GetFilterFuncIRI() (string, error) {
	id := f.GetID()
	if id == "" {
		return "", fmt.Errorf("Can't get FuncIRI for filter. ID is missing")
	}
	return fmt.Sprintf("url(#%s)", id), nil
}

func (f Filter) MarshalXMLAttr(n xml.Name) (xml.Attr, error) {
	funcIRI, err := f.GetFilterFuncIRI()
	return xml.Attr{Name: n, Value: funcIRI}, err

}

func (f *Filter) AddPrimitive(fp ...FilterPrimitive) {
	f.Primitives = append(f.Primitives, fp...)
}

type FilterSources string

func (f FilterSources) GetOut() (string, error) {
	return string(f), nil
}

func (f FilterSources) MarshalXMLAttr(n xml.Name) (xml.Attr, error) {
	out, _ := f.GetOut()
	return xml.Attr{Name: n, Value: out}, nil
}

const (
	SourceGraphic   FilterSources = "SourceGraphic"
	SourceAlpha     FilterSources = "SourceAlpha"
	BackgroundImage FilterSources = "BackgroundImage"
	BackgroundAlpha FilterSources = "BackgroundAlpha"
	FillPaint       FilterSources = "FillPaint"
	StrokePaint     FilterSources = "StrokePaint"
)
