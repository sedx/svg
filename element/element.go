package element

// Element is fake interface to filter elements
type Element interface {
	element()
}

// SVGElement used for composition to realise Element interface
type SVGElement struct{}

func (s SVGElement) element() {}

// SVGContainer is holder of other SVG elements
type SVGContainer struct {
	Elements []Element
}

// Add append element to elements slice
func (s *SVGContainer) Add(e Element) {
	s.Elements = append(s.Elements, e)
}
