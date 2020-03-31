package print

import (
	"fmt"

	"sewik/xml/spec"
)

func Attributes(attrs spec.Attributes) {
	for ak, av := range attrs.Get() {
		fmt.Printf(` %s="%d"`, ak, av)
	}
}
