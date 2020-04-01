package stats

import (
	"fmt"
	"strings"
)

func Print(e Elements) {
	printElements(e, 0, 0)
}

func printElements(e Elements, l, c int) {
	for name, elem := range e.Get() {
		fmt.Printf("\n%s<%s", strings.Repeat("  ", l), name)

		fmt.Printf(` _count="%d"`, elem.Cn)

		if elem.Cn < c {
			fmt.Print(` _optional="true"`)
		}

		printAttributes(elem.At)

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

func printAttributes(a Attributes) {
	for name, attr := range a.Get() {
		fmt.Printf(` %s="%d"`, name, attr)
	}
}
