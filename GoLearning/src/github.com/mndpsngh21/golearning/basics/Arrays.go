package basics

import "fmt"

func ArrayOps() {
	basicArray()
	stringArray()
	sliceOps()

}

func stringArray() {
	var students [3]string
	students[0] = "Mandeep"
	students[1] = "Sandeep"
	students[2] = "Jashan"
	fmt.Printf("Total Students: %v \n", len(students))
	fmt.Printf("Students: %v ", students)

}

func basicArray() {
	grades := [...]int{2, 5, 7}
	fmt.Printf("Grades: %v \n", grades)

}

func sliceOps() {
	grades := [...]int{2, 5, 7}
	n := grades[:2]
	fmt.Printf("Sliced Values :: %v \n", n)
	// slice out from and till index
	n1 := grades[:1:2]
	fmt.Printf("Sliced Values :: %v \n", n1)
	// create slice
	a := make([]int, 100)
	a = append(a, 33)
	fmt.Printf("Sliced Length:: %v \n", len(a))
	fmt.Printf("Sliced Cap:: %v \n", cap(a))

}
