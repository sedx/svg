package transformation

import (
	"fmt"
)

type Translate struct {
	Tx float64
	Ty float64
}

func (t Translate) GetTransform() string {
	return fmt.Sprintf("translate(%v,%v)", t.Tx, t.Ty)
}
