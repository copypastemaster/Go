package strings

import (
	"fmt"
	"strings"
	"unsafe"
)

func Stringers() {
	s := "abc"
	clone := strings.Clone(s)
	fmt.Println(s == clone)
	fmt.Println(unsafe.StringData(s) == unsafe.StringData(clone))
}
