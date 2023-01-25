package basics

import (
	"fmt"
	"log"
)

func FlowExamples() {
	fmt.Println(" **************** Flow examples ********************")
	fmt.Println("Start")
	panicTest()
	fmt.Println("End")
}

func panicTest() {
	fmt.Println("Executing Panic Test")
	// defer allows code to execute before returning from method, there are cases when it will helpful, such as file operations
	defer func() { // anonymous function
		// recover will help to continue flow
		if err := recover(); err != nil {
			log.Println("Error :", err)
		}
	}()
	// syntax for panic in some situation, similar to throw exception
	panic("I just want to panic")
	fmt.Println("Executing Panic Test Done")

}
