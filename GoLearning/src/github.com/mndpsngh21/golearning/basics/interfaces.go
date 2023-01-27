package basics

import "fmt"

func InterfaceOps() {
	var c Coder = Coder{}
	speak(&c)
	walk(&c)
}

type WorkingPerson interface {
	walk()
}

type Human interface {
	speak()
	WorkingPerson
}

type Coder struct {
}

// defining method for coder struct
func speak(c *Coder) {
	fmt.Println("Hey! I am also a human. I can speak")
}

func walk(c *Coder) {
	fmt.Println("Hey! I can walk")

}
