package main

import (
	"fmt"
	"golang-async/modules"
)

func main() {
	modules.NormalCallFunctions()

	fmt.Println("")

	modules.AsyncCallFunctions()
}
