package basics

import "fmt"

func MapOps() {

	mapOperations()
	structOPs()
}

func mapOperations() {
	statePopulation := map[string]int{
		"PB": 534343,
		"HR": 134343,
		"RJ": 32146,
	}
	// get population
	fmt.Printf("Population of HR %v \n", statePopulation["HR"])
	// another way to create map
	s := make(map[string]int)
	// value can be added to map
	s["HR"] = 8989
	fmt.Printf("Population of HR %v \n", s["HR"])
	// delete
	delete(s, "HR")
	_, isPresent := s["HR"]
	if !isPresent {
		fmt.Println("Value is not present in map")
	}

}

// package struct
type Employee struct {
	name string
	id   int
}

// we can have corelation
type Manager struct {
	Employee
	Code int
}

func structOPs() {

	emp1 := &Employee{name: "Mandeep", id: 1}
	fmt.Printf("Emp Name:: %v \n", emp1.name)

	// anynoums struct
	emp2 := struct{ name string }{name: "A"}
	fmt.Printf("Emp Name:: %v \n", emp2.name)

	// make relationship with composition

	mgr := &Manager{}
	mgr.name = "Mandeep"
	mgr.Code = 3344
	fmt.Printf("Manager Name:: %v \n", mgr)

}
