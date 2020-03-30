package print

import (
	"fmt"

	"sewik/xml/structure"
)

func Attributes(attrs structure.Attributes) {
	for ak, av := range attrs.Int() {
		fmt.Printf(` %s="%d"`, ak, av)
	}
}
