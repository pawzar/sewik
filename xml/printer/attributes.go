package printer

import (
	"fmt"

	"sewik/xml/spec"
)

func PrintAttributes(a spec.Attributes) {
	for name, attr := range a.Get() {
		fmt.Printf(` %s="%d"`, name, attr)
	}
}
