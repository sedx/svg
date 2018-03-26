package transformation

import (
	"fmt"

	"github.com/sedx/svg/types"
)

type Translate struct {
	Tx types.Length
	Ty types.Length
}

func (t Translate) GetTransform() string {
	return fmt.Sprintf("translate(%s,%s)", t.Tx, t.Ty)
}
