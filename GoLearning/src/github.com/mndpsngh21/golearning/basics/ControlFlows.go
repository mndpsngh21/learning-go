package basics

import "fmt"

func ControlFlow() {
	ifElseOps()
	switchOps()

}

func ifElseOps() {
	n := 6
	i := 6
	if n > 1 && i < 10 {
		fmt.Println("We are in range")
	} else if i == 1 && n == 1 {
		fmt.Println("Same value")

	} else {
		fmt.Println("We are not in range")
	}
}

func switchOps() {
	switch 4 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	// can have multiple case
	case 4, 5:
		fmt.Println("Four or Five")
	}

}
