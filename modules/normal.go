package modules

import (
	"fmt"
	"time"
)

func NormalCallFunctions() {
	fmt.Println("-------------- NORMAL CALL --------------")

	start := time.Now()

	func1()

	func2()

	func3()

	fmt.Printf("*** Time Duration : %s *** \n", time.Since(start))
}

func func1() {
	time.Sleep(170 * time.Millisecond)

	fmt.Println("func 1 done")
}

func func2() {
	time.Sleep(100 * time.Millisecond)

	fmt.Println("func 2 done")
}

func func3() {
	time.Sleep(300 * time.Millisecond)

	fmt.Println("func 3 done")
}
