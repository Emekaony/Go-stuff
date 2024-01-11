package using_stringers

import (
	"fmt"
	"strings"
)

func StringFields() {
	names := "emeka kamsi chelsea"
	arr := strings.Fields(names)
	fmt.Println(arr)
}
