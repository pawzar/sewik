package dom

import (
	"fmt"
	"strings"
)

func PrintDocument(doc *Document) {
	fmt.Println(doc.ProcInst)

	for _, v := range doc.Directives {
		fmt.Println(v)
	}

	PrintElement(doc.Root, 0)

	fmt.Println()
}

func PrintElement(e *Element, l int) {
	fmt.Printf("%s<%s", strings.Repeat("  ", l), e.Name)
	for _, a := range e.Attributes {
		fmt.Printf(` %s="%s"`, a.Name, a.Value)
	}

	if len(e.Children) > 0 {
		fmt.Printf(">%s", e.Text)
		for _, c := range e.Children {
			fmt.Printf("\n")
			PrintElement(c, l+1)
		}
		fmt.Printf("\n%s</%s>", strings.Repeat("  ", l), e.Name)
	} else if e.Text != "" {
		fmt.Printf(">%s</%s>", e.Text, e.Name)
	} else {
		fmt.Printf(" />")
	}
}
