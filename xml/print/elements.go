package print

import (
	"fmt"
	"strings"

	"sewik/xml/structure"
)

func Elements(elems structure.Elements, l, c int) {
	for k, elem := range elems.Int() {
		fmt.Printf("\n%s<%s", strings.Repeat("  ", l), k)

		fmt.Printf(` _count="%d"`, elem.Cn)

		if elem.Cn < c {
			fmt.Print(` _optional="true"`)
		}

		Attributes(elem.At)

		if elem.El.Len() > 0 {
			fmt.Print(">")

			Elements(elem.El, l+1, elem.Cn)

			fmt.Printf("\n%s</%s>", strings.Repeat("  ", l), k)
		} else {
			fmt.Print(" />")
		}
	}
}
