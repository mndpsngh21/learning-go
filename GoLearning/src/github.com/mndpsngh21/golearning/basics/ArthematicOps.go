package basics

import "fmt"

func ArthemicOPs() {
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Executing all Arthematic OPs")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
	add()
	subtract()
	divide()
	multiply()
	bitwiseOps()
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
}

func bitwiseOps() {
	i := 3
	j := 2
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Bitwise OPs")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Printf("%v&%v  = %v", i, j, j&i)
	fmt.Println()
	fmt.Printf("%v|%v  = %v", j, i, j|i)
	fmt.Println()
	fmt.Printf("%v^%v  = %v", j, i, j^i)
	fmt.Println()
	fmt.Printf("%v&^%v = %v", j, i, j&i)
	fmt.Println()
}

func multiply() {
	i := 3.5
	j := 4.6
	fmt.Printf("Multiplication of %v and %v is %v", i, j, (i * j))
	fmt.Println()

}

func divide() {
	i := 3
	j := 4
	fmt.Printf("Division of %v from %v is %v", i, j, (i / j))
	fmt.Println()
}

func subtract() {
	i := 3
	j := 4
	fmt.Printf("Subtraction of %v from %v is %v", i, j, (j - i))
	fmt.Println()
}

func add() {
	i := 3
	j := 4
	fmt.Printf("Sum of %v and %v is %v", i, j, (i + j))
	fmt.Println()
}
