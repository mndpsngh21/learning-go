package basics

import "fmt"

func PointerExample() {
	fmt.Println("********* Pointer concept ***********")
	base1()
	exmp2()
	functionCallExample()
}

func base1() {
	// understand reference context
	a := 1
	b := a
	fmt.Println(a, b)
	a = 5
	fmt.Println(a, b)
	c := &a
	fmt.Println(a, c)
	a = 67
	fmt.Println(a, c)
}

func exmp2() {
	var a int = 43
	var b *int = &a
	fmt.Println(a, b)
	fmt.Println(a, *b)

}

/*
*

	By default while calling functions values are getting copied and passed to function, we should avoid it by passing pointer reference


	func n(m string){

	}

	// call by value
	a:="Mandeep"
	n(a)

// call by pointer

func n(m *string){

	}

*
*/
func functionCallExample() {
	n := "Mandeep"
	callByValue(n)
	fmt.Printf("value after update %v \n", n)
	callByPointer(&n)
	fmt.Printf("value after update %v \n", n)

	v := returnValue()
	fmt.Printf("value from return  %v \n", v)

	r := returnPointer()
	fmt.Printf("value from return address = %v, value =%v \n", r, *r)

}

func callByValue(n string) {
	fmt.Println(n)
	n = "Someone else"
}

func callByPointer(n *string) {
	fmt.Println(n)
	*n = "Someone else"
}

// return from function

func returnValue() int {
	a := 1
	b := 5
	c := a + b
	return c
}

func returnPointer() *int {
	a := 1
	b := 5
	c := a + b
	return &c
}
