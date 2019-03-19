package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94800,
		},
	}
	// &: give me the memory address of the value(jim)
	// jimPointer := &jim 
	// jimPointer.updateName("Jimmy")

	jim.updateName("Jimmy")
	jim.print()
}


func (p *person) updateName(newFirstName string) { 
	(*p).firstName = newFirstName // *: give me the value of this memory address
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
