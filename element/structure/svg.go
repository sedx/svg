package structure

import (
	"encoding/xml"
	"io"
	"net/http"

	"github.com/sedx/svg/attributes"
	"github.com/sedx/svg/element"
	"github.com/sedx/svg/types"
)

// SVG is svg document
type SVG struct {
	element.SVGElement
	attributes.Measured
	attributes.ViewBoxed
	element.SVGContainer
	X       types.Coordinate `xml:"x,attr,omitempty"`
	Y       types.Coordinate `xml:"y,attr,omitempty"`
	XMLName xml.Name         `xml:"svg"`
}

type svgForMarshal struct {
	*SVG
	Version string `xml:"version,attr"`
	Ns      string `xml:"xmlns,attr"`
	Link    string `xml:"xmlns:xlink,attr"`
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
		Version: "1.2",
		Ns:      "http://www.w3.org/2000/svg",
		Link:    "http://www.w3.org/1999/xlink",
	}
}

func (s *SVG) WriteToResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "image/svg+xml")
	b, err := s.MarshalIndent()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
	}
	w.Write(b)
}
