package basics

import "fmt"

func VariableOPs() {
	allInts()

}

// we have signed usigned ints
// floats are floats32 float64

func allInts() {

	var i int = 1
	j := 4.5
	z := 7
	y := 189389232893
	b := true
	fmt.Printf("Int variable %v  type of %T ,%v  type of %T,%v type of %T,%v type of %T", i, i, j, j, z, z, y, y)
	fmt.Println()
	fmt.Printf("boolean %v type %T", b, b)
	fmt.Println()

	var s string = "mandeep"
	var r rune = 'D' // similar to char

	fmt.Printf(" %v %T", s, s)
	fmt.Println()
	fmt.Printf(" %v %T", r, r)
	fmt.Println()
}
