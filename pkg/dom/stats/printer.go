package stats

import (
	"fmt"
	"strings"
)

func PrintXML(e Elements) {
	printElements(e, 0, 0)
	fmt.Println()
}

func printElements(e Elements, l, c int) {
	for name, elem := range e.Get() {
		fmt.Printf("\n%s<%s", strings.Repeat("  ", l), name)

		printAttributes(elem.At)

		if elem.El.Len() > 0 {
			fmt.Print(">")

			printElements(elem.El, l+1, elem.Cn)

			fmt.Printf("\n%s</%s>", strings.Repeat("  ", l), name)
		} else {
			fmt.Print(" />")
		}
	}
}

func printAttributes(a Attributes) {
	for name, attr := range a.Get() {
		fmt.Printf(` %s="%d"`, name, attr)
	}
}
