package models2

import (
	"fmt"
)

func debug(format string, a ...interface{}) {
	// if !strings.Contains(format, "setBson") {
		return
	// }
	fmt.Print("[serialisation] ")
	fmt.Printf(format, a...)
	fmt.Println()
}
