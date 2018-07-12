package server

import (
	"fmt"
)

func debug(format string, a ...interface{}) {
	// if !strings.Contains(format, "setBson") {
		return
	// }
	fmt.Print("[server] ")
	fmt.Printf(format, a...)
	fmt.Println()
}