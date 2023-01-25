package main

import (
	"fmt"
	"strconv"

	OPs "github.com/mndpsngh21/golearning/basics"
)

var s string = "545" // package level scope

func main() {
	var m int = 5
	i := 45 // local scope
	d, _error := strconv.Atoi(s)

	if _error == nil {
		fmt.Printf("local scope %v, package scope %v, local 2 version %v  ", i, d, m)
		fmt.Println()
	}
	OPs.ArthemicOPs()
	OPs.VariableOPs()
	OPs.ConstantOPs()
	breaker()
	OPs.ArrayOps()
	breaker()
	OPs.MapOps()
	breaker()
	OPs.ControlFlow()
	breaker()
	OPs.FlowExamples()
	breaker()
	OPs.PointerExample()
}

func breaker() {
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++")
}
