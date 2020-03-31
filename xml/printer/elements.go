package printer

import (
	"fmt"
	"strings"

	"sewik/xml/spec"
)

func PrintElements(e spec.Elements) {
	printElements(e, 0, 0)
}

func printElements(e spec.Elements, l, c int) {
	for name, elem := range e.Get() {
		fmt.Printf("\n%s<%s", strings.Repeat("  ", l), name)

		fmt.Printf(` _count="%d"`, elem.Cn)

		if elem.Cn < c {
			fmt.Print(` _optional="true"`)
		}

		PrintAttributes(elem.At)

		if elem.El.Len() > 0 {
			fmt.Print(">")

			printElements(elem.El, l+1, elem.Cn)

			fmt.Printf("\n%s</%s>", strings.Repeat("  ", l), name)
		} else {
			fmt.Print(" />")
		}
	}

	fmt.Println()
}
